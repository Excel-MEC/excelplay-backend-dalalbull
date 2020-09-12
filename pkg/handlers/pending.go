package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/database"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/env"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/httperrors"
	"github.com/dgrijalva/jwt-go"
)

// Pending returns information about pending stocks
func Pending(db *database.DB, env *env.Config) httperrors.Handler {
	return func(w http.ResponseWriter, r *http.Request) *httperrors.HTTPError {
		props, _ := r.Context().Value("props").(jwt.MapClaims)
		uid := props["sub"].(string)

		var pending []database.PendingData
		err := db.GetPendingStocks(uid, &pending)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Could not retrieve pending data", http.StatusInternalServerError}
		}

		jsonRes, err := json.Marshal(pending)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Could not serialize json", http.StatusInternalServerError}
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonRes)
		return nil
	}
}
