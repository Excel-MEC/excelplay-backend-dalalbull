package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/utils"

	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/database"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/env"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/httperrors"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/strconst"
	"github.com/dgrijalva/jwt-go"
)

// Cancel a pending transaction
func Cancel(db *database.DB, env env.Config) httperrors.Handler {
	type pendingStockToDelete struct {
		Company string `json:"company"`
	}
	type result struct {
		Message string `json:"msg"`
	}
	return func(w http.ResponseWriter, r *http.Request) *httperrors.HTTPError {
		var msg string

		cclose := utils.IsMarketOpen()
		if !cclose {
			msg = "Markets are closed"
		} else {
			props, _ := r.Context().Value("props").(jwt.MapClaims)
			uid := props["sub"].(string)
			var p pendingStockToDelete
			err := json.NewDecoder(r.Body).Decode(&p)
			if err != nil {
				return &httperrors.HTTPError{r, err, strconst.InvalidJSON, http.StatusBadRequest}
			}

			_, err = db.DeletePending(uid, p.Company)
			if err != nil {
				return &httperrors.HTTPError{r, err, "Could not delete pending transaction", http.StatusInternalServerError}
			}
			msg = "Specified pending order has been cancelled"
		}

		jsonRes, err := json.Marshal(result{Message: msg})
		if err != nil {
			return &httperrors.HTTPError{r, err, strconst.JSONEncodingFail, http.StatusInternalServerError}
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonRes)
		return nil
	}
}
