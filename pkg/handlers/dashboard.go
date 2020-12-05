package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/database"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/env"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/httperrors"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/strconst"
	"github.com/dgrijalva/jwt-go"
)

// GetDashboard gets the details that are displayed on a user's dashboard
func GetDashboard(db *database.DB, env *env.Config) httperrors.Handler {
	return func(w http.ResponseWriter, r *http.Request) *httperrors.HTTPError {
		props, _ := r.Context().Value("props").(jwt.MapClaims)
		userID, _ := strconv.Atoi(props["user_id"].(string))

		var userCount int
		err := db.GetTotalUsers(&userCount)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Could not get total users", http.StatusInternalServerError}
		}

		var stockHoldingsBuy []database.Stock
		err = db.GetStockHoldingsBuy(userID, &stockHoldingsBuy)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Could not get bought stocks data", http.StatusInternalServerError}
		}
		var stockHoldingsShortSell []database.Stock
		err = db.GetStockHoldingsShortSell(userID, &stockHoldingsShortSell)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Could not get shorted stocks data", http.StatusInternalServerError}
		}
		// Combine info on bought and shorted stocks into a single array
		for _, v := range stockHoldingsShortSell {
			stockHoldingsBuy = append(stockHoldingsBuy, v)
		}

		var topGainers []database.StockChange
		err = db.GetTopGainers(&topGainers)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Could not get top gainers", http.StatusInternalServerError}
		}

		var topLosers []database.StockChange
		err = db.GetTopGainers(&topLosers)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Could not get top losers", http.StatusInternalServerError}
		}

		var mostActiveVol []database.StockVolume
		err = db.GetTopVol(&mostActiveVol)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Could not get stocks with top volume", http.StatusInternalServerError}
		}

		var mostActiveVal []database.StockValue
		err = db.GetTopVal(&mostActiveVal)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Could not get stocks with top value", http.StatusInternalServerError}
		}

		jsonRes, err := json.Marshal(database.DashboardData{
			TotalUsers:    userCount,
			StockHoldings: stockHoldingsBuy,
			TopGainers:    topGainers,
			TopLosers:     topLosers,
			MostActiveVol: mostActiveVol,
			MostActiveVal: mostActiveVal,
		})
		if err != nil {
			return &httperrors.HTTPError{r, err, strconst.JSONEncodingFail, http.StatusInternalServerError}
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonRes)
		return nil
	}
}
