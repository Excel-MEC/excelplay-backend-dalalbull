package handlers

import (
	"net/http"

	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/httperrors"
)

// HandleAdmin handles any requests to the /api/admin endpoint
func HandleAdmin() httperrors.Handler {
	return func(w http.ResponseWriter, r *http.Request) *httperrors.HTTPError {
		w.Write([]byte("Admin"))
		return nil
	}
}
