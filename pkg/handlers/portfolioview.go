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

// PortfolioView rturns portfolio of a single user
func PortfolioView(db *database.DB, env *env.Config) httperrors.Handler {
	return func(w http.ResponseWriter, r *http.Request) *httperrors.HTTPError {
		props, _ := r.Context().Value("props").(jwt.MapClaims)
		userID := strconv.Atoi(props["user_id"].(string))
		// TODO: Remove completely if totalUserCount is not being used in the frontend
		// var totalUserCount int
		// err := db.GetTotalUsers(&totalUserCount)
		// if err != nil {
		// 	return &httperrors.HTTPError{r, err, "Could not get total user count", http.StatusInternalServerError}
		// }
		var portfolio database.Portfolio
		err := db.GetPortfolio(&portfolio, userID)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Could not get portfolio", http.StatusInternalServerError}
		}
		jsonRes, err := json.Marshal(portfolio)
		if err != nil {
			return &httperrors.HTTPError{r, err, strconst.JSONEncodingFail, http.StatusInternalServerError}
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonRes)
		return nil
	}
}
