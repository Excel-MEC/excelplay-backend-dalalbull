package routes

import (
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/database"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/env"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/handlers"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/middlewares"
	"github.com/gorilla/mux"
)

// Router wraps mux.Router to add route init method
type Router struct {
	*mux.Router
}

// NewRouter setups and returns a new router
func NewRouter() *Router {
	return &Router{
		mux.NewRouter(),
	}
}

// Routes initializes the routes of the api
func (router *Router) Routes(db *database.DB, config *env.Config) {
	router.Handle("/admin/", middlewares.ErrorsMiddleware(handlers.HandleAdmin())).Methods("GET")
	router.Handle("/api/ping", middlewares.ErrorsMiddleware(handlers.HeartBeat())).Methods("GET")
	router.Handle("/api/question",
		middlewares.ErrorsMiddleware(
			middlewares.AuthMiddleware(
				handlers.HandleNextQuestion(db, config),
				config,
			),
		),
	).Methods("GET")
	router.Handle("/api/submit",
		middlewares.ErrorsMiddleware(
			middlewares.AuthMiddleware(
				handlers.HandleSubmission(db, config),
				config,
			),
		),
	).Methods("POST")

	// LoggerMiddleware does not have to be selectively applied because it applies to all endpoints
	router.Use(middlewares.LoggerMiddleware)
}

// Note that ErrorsMiddleware must always be the outermost middleware of the selectively applied middlewares
// as it handles errors from all internal functions and is the only funcion in the chain that returns the
// http.Handler which is expected by router.Handle in it's 2nd arg.
