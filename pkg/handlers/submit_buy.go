package handlers

import (
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/utils"
	"encoding/json"
	"net/http"

	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/strconst"

	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/database"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/env"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/httperrors"
)

// GetCompanyDetails gets all the stock details of a single company
func SubmitBuy(db *database.DB, env *env.Config) httperrors.Handler {
	type TransactionType string
	const (
		buy TransactionType = iota
		sell TransactionType
	)
	type request struct {
		Quantity int,
		Company string,
		Pending float32, // empty if not pending
		Type TransactionType, // buy or shortsell
	}
	type response struct {
		Msg string
	}
	func buyTransaction(w http.ResponseWriter, req request) *httperrors.HTTPError {
		var companyInfo database.CompanyInfo
		err = db.GetCompanyStockInfo(req.Company, &companyInfo)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Incorrect company name", http.StatusBadRequest}
		}
		currPrice := companyInfo.CurrPrice

		props, _ := r.Context().Value("props").(jwt.MapClaims)
		userID := props["sub"].(string)
		var portfolio database.Portfolio
		err := db.GetPortfolio(&portfolio, userID)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Could not retrieve user portfolio", http.StatusBadRequest}
		}
		noOfTrans := portfolio.NoTrans
		margin := portfolio.Margin

		pending := req.Pending
		if pending != nil {
			if pending == curr_price {
				return &httperrors.HTTPError{r, nil, "Pending price cannot be equal to current price", http.StatusBadRequest}
			}
			jsonRes, err := json.Marshal(companyInfo)
			if err != nil {
				return &httperrors.HTTPError{r, err, "Could not serialize json", http.StatusInternalServerError}
			}

			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonRes)
			return nil
		} 

		brokerage := utils.CalculateBrokerage(noOfTrans, req.Quantity, currPrice)

		userCashBal := portfolio.CashBal
		if userCashBal - (req.Quantity * currPrice) - brokerage < 0 {
			return &httperrors.HTTPError{r, nil, "Not enough balance", http.StatusBadRequest}
		}

		txn := database.TransactionBuy{}
		err = db.Get(&txn, "select * from transaction_buy where user_id = ? and symbol = ?", userID, req.Symbol)
		if err != nil {
				
		}

	}
	func shortSellTransaction(w http.ResponseWriter, req request) *httperrors.HTTPError {
		var companyInfo database.CompanyInfo
		err = db.GetCompanyStockInfo(req.Company, &companyInfo)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Incorrect company name", http.StatusBadRequest}
		}

		pending := req.Pending
		if pending == nil {
			jsonRes, err := json.Marshal(companyInfo)
			if err != nil {
				return &httperrors.HTTPError{r, err, "Could not serialize json", http.StatusInternalServerError}
			}

			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonRes)
			return nil
		}
		} else {

		}
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
			return buyTransaction(w, req)
		}
		return shortSellTransaction(w, req)

	}
}