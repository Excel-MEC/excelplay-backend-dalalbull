package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/database"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/env"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/httperrors"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/strconst"
)

// Ticker returns details of all stocks to display in a ticker
func Ticker(db *database.DB, env *env.Config) httperrors.Handler {
	return func(w http.ResponseWriter, r *http.Request) *httperrors.HTTPError {
		var tickerData []database.TickerData
		err := db.GetTickerData(&tickerData)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Could not get ticker data", http.StatusInternalServerError}
		}
		jsonRes, err := json.Marshal(tickerData)
		if err != nil {
			return &httperrors.HTTPError{r, err, strconst.JSONEncodingFail, http.StatusInternalServerError}
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonRes)
		return nil
	}
}
