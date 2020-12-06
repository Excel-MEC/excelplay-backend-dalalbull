package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/utils"
	"github.com/dgrijalva/jwt-go"

	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/strconst"

	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/database"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/env"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/httperrors"
)

// GetCompanyDetails gets all the stock details of a single company
func SubmitBuy(db *database.DB, env *env.Config) httperrors.Handler {
	type TransactionType int
	const (
		buy TransactionType = iota
		sell
	)
	type request struct {
		Quantity int
		Company  string
		Pending  float32         // empty if not pending
		Type     TransactionType // buy or shortsell
	}
	type response struct {
		Msg string
	}
	buyTransaction := func(w http.ResponseWriter, r *http.Request, req request) *httperrors.HTTPError {
		var companyInfo database.CompanyInfo
		err := db.GetCompanyStockInfo(req.Company, &companyInfo)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Incorrect company name", http.StatusBadRequest}
		}
		jsonRes, err := json.Marshal(companyInfo)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Could not serialize json", http.StatusInternalServerError}
		}
		currPrice := companyInfo.CurrPrice

		props, _ := r.Context().Value("props").(jwt.MapClaims)
		userID, _ := strconv.Atoi(props["user_id"].(string))
		var portfolio database.Portfolio
		err = db.GetPortfolio(&portfolio, userID)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Could not retrieve user portfolio", http.StatusBadRequest}
		}
		noOfTrans := portfolio.NoTrans
		// margin := portfolio.Margin

		pending := req.Pending
		if pending != 0 {
			if pending == currPrice {
				return &httperrors.HTTPError{r, nil, "Pending price cannot be equal to current price", http.StatusBadRequest}
			}

			// TODO : Create a pending txn

			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonRes)
			return nil
		}

		brokerage := utils.CalculateBrokerage(noOfTrans, req.Quantity, currPrice)

		userCashBal := portfolio.CashBal
		if userCashBal-(float32(req.Quantity)*currPrice)-brokerage < 0 {
			return &httperrors.HTTPError{r, nil, "Not enough balance", http.StatusBadRequest}
		}

		txn := database.TransactionBuy{}
		err = db.GetTransactionBuy(userID, req.Company, &txn)
		if err != nil {
			// Create new txn buy
			err = db.CreateNewTransactionBuy(userID, req.Company, req.Quantity, currPrice)
			if err != nil {
				return &httperrors.HTTPError{r, err, "Internal server error", http.StatusInternalServerError}
			}

		} else {
			// Update existing txn buy
			val := (currPrice*float32(req.Quantity) + txn.Value*float32(txn.Quantity)) / float32(req.Quantity+txn.Quantity)
			err = db.UpdateTransactionBuy(userID, req.Company, req.Quantity+txn.Quantity, val)
			if err != nil {
				return &httperrors.HTTPError{r, err, "Internal server error", http.StatusInternalServerError}
			}
		}

		err = db.UpdatePortfolio(userID, portfolio.CashBal-brokerage, portfolio.NoTrans+1, portfolio.Margin)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Internal server error", http.StatusInternalServerError}
		}

		// Create new history entry
		history := database.History{
			Symbol:   req.Company,
			BuySS:    int(buy),
			Quantity: float32(req.Quantity),
			Price:    currPrice,
		}
		err = db.CreateNewHistory(userID, history)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Internal server error", http.StatusInternalServerError}
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonRes)
		return nil
	}
	shortSellTransaction := func(w http.ResponseWriter, r *http.Request, req request) *httperrors.HTTPError {
		var companyInfo database.CompanyInfo
		err := db.GetCompanyStockInfo(req.Company, &companyInfo)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Incorrect company name", http.StatusBadRequest}
		}
		jsonRes, err := json.Marshal(companyInfo)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Could not serialize json", http.StatusInternalServerError}
		}
		currPrice := companyInfo.CurrPrice

		props, _ := r.Context().Value("props").(jwt.MapClaims)
		userID, _ := strconv.Atoi(props["user_id"].(string))
		var portfolio database.Portfolio
		err = db.GetPortfolio(&portfolio, userID)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Could not retrieve user portfolio", http.StatusBadRequest}
		}
		noOfTrans := portfolio.NoTrans

		pending := req.Pending
		if pending == 0 {
			jsonRes, err := json.Marshal(companyInfo)
			if err != nil {
				return &httperrors.HTTPError{r, err, "Could not serialize json", http.StatusInternalServerError}
			}

			// TODO : Create a pending txn

			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonRes)
			return nil
		}
		brokerage := utils.CalculateBrokerage(noOfTrans, req.Quantity, currPrice)

		userCashBal := portfolio.CashBal
		if userCashBal-(float32(req.Quantity)*currPrice)/2-brokerage < 0 {
			return &httperrors.HTTPError{r, nil, "Not enough balance", http.StatusBadRequest}
		}

		txn := database.TransactionShortSell{}
		err = db.GetTransactionShortSell(userID, req.Company, &txn)
		if err != nil {
			// Create new txn shortsell
			err = db.CreateNewTransactionShortSell(userID, req.Company, req.Quantity, currPrice)
			if err != nil {
				return &httperrors.HTTPError{r, err, "Internal server error", http.StatusInternalServerError}
			}

		} else {
			// Update existing txn buy
			val := (currPrice*float32(req.Quantity) + txn.Value*float32(txn.Quantity)) / float32(req.Quantity+txn.Quantity)
			err = db.UpdateTransactionShortSell(userID, req.Company, req.Quantity+txn.Quantity, val)
			if err != nil {
				return &httperrors.HTTPError{r, err, "Internal server error", http.StatusInternalServerError}
			}
		}

		margin := float32(req.Quantity) * currPrice / 2
		err = db.UpdatePortfolio(userID, portfolio.CashBal-brokerage, portfolio.NoTrans+1, portfolio.Margin+margin)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Internal server error", http.StatusInternalServerError}
		}

		// Create new history entry
		history := database.History{
			Symbol:   req.Company,
			BuySS:    int(sell),
			Quantity: float32(req.Quantity),
			Price:    currPrice,
		}
		err = db.CreateNewHistory(userID, history)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Internal server error", http.StatusInternalServerError}
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonRes)
		return nil
	}
	return func(w http.ResponseWriter, r *http.Request) *httperrors.HTTPError {
		var req request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return &httperrors.HTTPError{r, err, strconst.InvalidJSON, http.StatusBadRequest}
		}
		if req.Quantity <= 0 {
			return &httperrors.HTTPError{r, err, "Quantity cannot be 0", http.StatusBadRequest}
		}

		marketOpen := utils.IsMarketOpen()
		if !marketOpen {
			return &httperrors.HTTPError{r, err, "Markets Closed", http.StatusBadRequest}
		}

		reqtype := req.Type
		if reqtype == buy {
			return buyTransaction(w, r, req)
		}
		return shortSellTransaction(w, r, req)

	}
}
