package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/database"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/env"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/httperrors"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/strconst"
)

// StockSymbol returns the symbols of all the companies
func StockSymbols(db *database.DB, env *env.Config) httperrors.Handler {
	type Companies struct {
		Companies []string `json:"companies"`
	}
	return func(w http.ResponseWriter, r *http.Request) *httperrors.HTTPError {
		var data []string
		err := db.GetCompanySymbols(&data)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Could not get company data", http.StatusInternalServerError}
		}

		jsonRes, err := json.Marshal(Companies{Companies: data})
		if err != nil {
			return &httperrors.HTTPError{r, err, strconst.JSONEncodingFail, http.StatusInternalServerError}
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonRes)
		return nil
	}
}
