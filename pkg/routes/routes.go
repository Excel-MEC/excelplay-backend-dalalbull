package routes

import (
	_ "github.com/Excel-MEC/excelplay-backend-dalalbull/cmd/excelplay-backend-dalalbull/docs"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/database"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/env"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/handlers"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/httperrors"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/middlewares"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
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
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	router.Handle("/api/ping", middlewares.ErrorsMiddleware(
		httperrors.Handler(
			handlers.HeartBeat(),
		),
	)).Methods("GET")
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
	).Methods("POST")
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
	router.Handle("/api/cancels", // API Change, what does 'iddel' in the input correspond to, since it isn't in the model?
		middlewares.ErrorsMiddleware(
			httperrors.Handler(
				middlewares.AuthMiddleware(
					handlers.Cancel(db, config),
					config,
				),
			),
		),
	).Methods("POST")
	router.Handle("/api/stockinfo",
		middlewares.ErrorsMiddleware(
			httperrors.Handler(
				middlewares.AuthMiddleware(
					handlers.StockSymbols(db, config),
					config,
				),
			),
		),
	).Methods("GET")
	router.Handle("/api/currentprice",
		middlewares.ErrorsMiddleware(
			httperrors.Handler(
				middlewares.AuthMiddleware(
					handlers.CurrentPrice(db, config),
					config,
				),
			),
		),
	).Methods("POST")
	router.Handle("/api/history",
		middlewares.ErrorsMiddleware(
			httperrors.Handler(
				middlewares.AuthMiddleware(
					handlers.History(db, config),
					config,
				),
			),
		),
	).Methods("GET")
	router.Handle("/api/submit_buy",
		middlewares.ErrorsMiddleware(
			httperrors.Handler(
				middlewares.AuthMiddleware(
					handlers.SubmitBuy(db, config),
					config,
				),
			),
		),
	).Methods("POST")
	router.Handle("/api/is_share_market_open",
		middlewares.ErrorsMiddleware(
			httperrors.Handler(
				middlewares.AuthMiddleware(
					handlers.IsShareMarketOpen(db, config),
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
