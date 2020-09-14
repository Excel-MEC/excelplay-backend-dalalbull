package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/strconst"

	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/database"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/env"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/httperrors"
)

// GetCompanyDetails gets all the stock details of a single company
func GetCompanyDetails(db *database.DB, env *env.Config) httperrors.Handler {
	type company struct {
		Symbol string `json:"symbol"`
	}
	return func(w http.ResponseWriter, r *http.Request) *httperrors.HTTPError {
		var c company
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			return &httperrors.HTTPError{r, err, strconst.InvalidJSON, http.StatusBadRequest}
		}

		var companyInfo database.CompanyInfo
		err = db.GetCompanyStockInfo(c.Symbol, &companyInfo)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Incorrect company name", http.StatusBadRequest}
		}

		jsonRes, err := json.Marshal(companyInfo)
		if err != nil {
			return &httperrors.HTTPError{r, err, strconst.JSONEncodingFail, http.StatusInternalServerError}
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonRes)
		return nil
	}
}
