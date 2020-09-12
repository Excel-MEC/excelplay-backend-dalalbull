package routes

import (
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/database"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/env"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/handlers"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/httperrors"
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
	router.HandleFunc("/admin/", handlers.HandleAdmin).Methods("GET")
	router.Handle("/api/handshake",
		middlewares.ErrorsMiddleware(
			httperrors.Handler(
				middlewares.AuthMiddleware(
					handlers.InitUser(db, config),
					config,
				),
			),
		),
	).Methods("GET")
	router.Handle("/api/portfolioview",
		middlewares.ErrorsMiddleware(
			httperrors.Handler(
				middlewares.AuthMiddleware(
					handlers.PortfolioView(db, config),
					config,
				),
			),
		),
	).Methods("GET")
	router.Handle("/api/ticker",
		middlewares.ErrorsMiddleware(
			httperrors.Handler(
				middlewares.AuthMiddleware(
					handlers.Ticker(db, config),
					config,
				),
			),
		),
	).Methods("GET")
	router.Handle("/api/leaderboard",
		middlewares.ErrorsMiddleware(
			httperrors.Handler(
				middlewares.AuthMiddleware(
					handlers.GetLeaderboard(db, config),
					config,
				),
			),
		),
	).Methods("GET")
	router.Handle("/api/dashboard",
		middlewares.ErrorsMiddleware(
			httperrors.Handler(
				middlewares.AuthMiddleware(
					handlers.GetDashboard(db, config),
					config,
				),
			),
		),
	).Methods("GET")
	router.Handle("/api/companydetails",
		middlewares.ErrorsMiddleware(
			httperrors.Handler(
				middlewares.AuthMiddleware(
					handlers.GetCompanyDetails(db, config),
					config,
				),
			),
		),
	).Methods("GET")
	router.Handle("/api/sell",
		middlewares.ErrorsMiddleware(
			httperrors.Handler(
				middlewares.AuthMiddleware(
					handlers.SellInfo(db, config),
					config,
				),
			),
		),
	).Methods("GET")
	router.Handle("/api/pending", // API Change
		middlewares.ErrorsMiddleware(
			httperrors.Handler(
				middlewares.AuthMiddleware(
					handlers.Pending(db, config),
					config,
				),
			),
		),
	).Methods("GET")

	// LoggerMiddleware does not have to be selectively applied because it applies to all endpoints
	router.Use(middlewares.LoggerMiddleware)
}

/*
Note that ErrorsMiddleware must always be the outermost middleware of the selectively applied middlewares
as it handles errors from all internal functions and is the only funcion in the chain that returns the
http.Handler which is expected by router.Handle in it's 2nd arg.
*/
