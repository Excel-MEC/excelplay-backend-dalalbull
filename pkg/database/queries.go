package database

import (
	"database/sql"
)

// CreateNewUser creates a new user, who starts at level 1.
func (db *DB) CreateNewUser(uuid string, name string) (sql.Result, error) {
	return db.Exec("insert into duser values($1,$2)", uuid, name)
}

// GetUser gets the details of a user
func (db *DB) GetUser(currUser *User, uuid string) error {
	return db.Get(currUser, "select name from duser where id = $1", uuid)
}

// GetTotalUsers returns the number of total users in the leaderboard
func (db *DB) GetTotalUsers(n *int) error {
	return db.Get(n, "select count(*) from duser")
}

// GetPortfolio returns the portfolio of a user
func (db *DB) GetPortfolio(portfolio *Portfolio, uuid string) error {
	return db.Get(portfolio, "select cash_bal, net_worth, rank, no_trans, margin from portfolio where user_id = $1", uuid)
}

// GetTickerData returns stock data for the ticker
func (db *DB) GetTickerData(tickerData *[]TickerData) error {
	return db.Select(tickerData, "select name, current_price, change_per from stocks_data")
}

// CreatePortfolio creates a portfolio for a new user
func (db *DB) CreatePortfolio(uuid string) (sql.Result, error) {
	var totalUsers int
	db.Get(totalUsers, "select count(*) from portfolio")
	return db.Exec("insert into portfolio (user_id, rank) values($1, $2)", uuid, totalUsers+1)
}

// GetTopGainers gets the top 5 stocks with largest gains
func (db *DB) GetTopGainers(topGainers *[]StockChange) error {
	return db.Select(topGainers, "select symbol, change_per from stocks_data order by change_per desc limit 5")
}

// GetTopLosers gets the top 5 stocks with largest losses
func (db *DB) GetTopLosers(topLosers *[]StockChange) error {
	return db.Select(topLosers, "select symbol, change_per from stocks_data order by change_per limit 5")
}

// GetTopVol gets the top 5 stocks with the highest trade quantity
func (db *DB) GetTopVol(mostActiveVol *[]StockVolume) error {
	return db.Select(mostActiveVol, "select symbol, trade_qty from stocks_data order by trade_qty desc limit 5")
}

// GetTopVal gets the top 5 stocks with the highest trade value
func (db *DB) GetTopVal(mostActiveVal *[]StockValue) error {
	return db.Select(mostActiveVal, "select symbol, trade_value from stocks_data order by trade_value desc limit 5")
}

// GetLeaderboard gets the users list in the descending order of level,
// and for users on the same level, in the ascending order of last submission time.
func (db *DB) GetLeaderboard(users *[]User) error {
	return db.Select(users, "select name, curr_level from duser order by curr_level desc, last_anstime")
}
