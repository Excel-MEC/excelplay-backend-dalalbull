package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/utils"

	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/database"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/env"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/httperrors"
	"github.com/dgrijalva/jwt-go"
)

// SellInfo returns information about all the stocks a user currently holds
func SellInfo(db *database.DB, env *env.Config) httperrors.Handler {
	return func(w http.ResponseWriter, r *http.Request) *httperrors.HTTPError {
		props, _ := r.Context().Value("props").(jwt.MapClaims)
		userID := props["sub"].(string)

		noStock := false
		transactions := make([]database.SellData, 0)
		var stocks []database.Stock
		marketIsOpen := utils.IsMarketOpen()
		if marketIsOpen {
			err := db.GetStockHoldingsBuy(userID, &stocks)
			if err != nil && err.Error() != "sql: no rows in result set" {
				return &httperrors.HTTPError{r, err, "Could not get bought stocks data", http.StatusInternalServerError}
			}
			for _, v := range stocks {
				transactions = append(transactions, database.SellData{
					Company:      v.Company,
					TypeOfTrade:  v.Type,
					ShareInHand:  v.Number,
					CurrentPrice: v.Current,
					Gain:         ((v.Purchase - v.Current) / v.Purchase) * 100,
					TypeOfTrans:  "SELL",
				})
			}
			err = db.GetStockHoldingsShortSell(userID, &stocks)
			if err != nil && err.Error() != "sql: no rows in result set" {
				return &httperrors.HTTPError{r, err, "Could not get bought stocks data", http.StatusInternalServerError}
			}
			for _, v := range stocks {
				transactions = append(transactions, database.SellData{
					Company:      v.Company,
					TypeOfTrade:  v.Type,
					ShareInHand:  v.Number,
					CurrentPrice: v.Current,
					Gain:         ((v.Purchase - v.Current) / v.Purchase) * 100,
					TypeOfTrans:  "SHORT COVER",
				})
			}
			if len(transactions) == 0 {
				noStock = true
			}
		}

		jsonRes, err := json.Marshal(database.SellDataWrapper{
			CClose:  !marketIsOpen,
			NoStock: noStock,
			Data:    transactions,
		})
		if err != nil {
			return &httperrors.HTTPError{r, err, "Could not serialize json", http.StatusInternalServerError}
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonRes)
		return nil
	}
}
