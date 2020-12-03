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

// CurrentPrice is called when user selects a company, to send back the current stock price of that company
func CurrentPrice(db *database.DB, env *env.Config) httperrors.Handler {
	type company struct {
		Company string `json:"company"`
	}
	return func(w http.ResponseWriter, r *http.Request) *httperrors.HTTPError {
		props, _ := r.Context().Value("props").(jwt.MapClaims)
		userID := strconv.Atoi(props["user_id"].(string))

		var c company
		var currPrice float32
		var currPriceInfo database.CurrentPriceInfo

		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			return &httperrors.HTTPError{r, err, strconst.InvalidJSON, http.StatusBadRequest}
		}

		err = db.GetCurrentPrice(c.Company, &currPrice)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Incorrect company name", http.StatusBadRequest}
		}
		err = db.GetUserInfoForCurrentPrice(userID, &currPriceInfo)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Failed to get user info", http.StatusInternalServerError}
		}
		currPriceInfo.CurrPrce = currPrice

		jsonRes, err := json.Marshal(currPriceInfo)
		if err != nil {
			return &httperrors.HTTPError{r, err, strconst.JSONEncodingFail, http.StatusInternalServerError}
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonRes)
		return nil
	}
}
