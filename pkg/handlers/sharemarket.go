package handlers

import (
	"net/http"

	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/utils"

	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/database"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/env"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/httperrors"
)

// IsShareMarketOpen is used to send a true or false value to the frontend about the status of the stock market
func IsShareMarketOpen(db *database.DB, env *env.Config) httperrors.Handler {
	return func(w http.ResponseWriter, r *http.Request) *httperrors.HTTPError {
		w.WriteHeader(http.StatusOK)
		if utils.IsMarketOpen() {
			w.Write([]byte("true"))
		} else {
			w.Write([]byte("false"))
		}
		return nil
	}
}
