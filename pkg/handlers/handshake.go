package handlers

import (
	"net/http"

	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/database"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/env"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/httperrors"
	"github.com/dgrijalva/jwt-go"
)

// InitUser handles any request made to the /api/question/ endpoint
func InitUser(db *database.DB, env *env.Config) httperrors.Handler {
	return func(w http.ResponseWriter, r *http.Request) *httperrors.HTTPError {
		// Obtain values from JWT
		props, _ := r.Context().Value("props").(jwt.MapClaims)
		userID := props["sub"].(string)
		name := props["name"].(string)

		var currUser database.User
		err := db.GetUser(&currUser, userID)
		if err != nil && err.Error() == "sql: no rows in result set" {
			_, err := db.CreateNewUser(userID, name)
			if err != nil {
				return &httperrors.HTTPError{r, err, "Could not create new user", http.StatusInternalServerError}
			}
		} else if err != nil {
			return &httperrors.HTTPError{r, err, "Invalid user", http.StatusInternalServerError}
		}

		var portfolio database.Portfolio
		err = db.GetPortfolio(&portfolio, userID)
		if err != nil && err.Error() == "sql: no rows in result set" {
			_, err := db.CreatePortfolio(userID)
			if err != nil {
				return &httperrors.HTTPError{r, err, "Could not create new portfolio", http.StatusInternalServerError}
			}
		} else if err != nil {
			return &httperrors.HTTPError{r, err, "Invalid portfolio", http.StatusInternalServerError}
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("success"))
		return nil
	}
}
