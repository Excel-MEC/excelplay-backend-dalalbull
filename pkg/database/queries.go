package database

import (
	"database/sql"
)

// CreateNewUser creates a new user, who starts at level 1.
func (db *DB) CreateNewUser(uuid int, name string) (sql.Result, error) {
	return db.Exec("insert into duser values($1,$2)", uuid, name)
}

// GetUser gets the details of a user
func (db *DB) GetUser(currUser *User, uuid int) error {
	return db.Get(currUser, "select name from duser where id = $1", uuid)
}

// GetTotalUsers returns the number of total users in the leaderboard
func (db *DB) GetTotalUsers(n *int) error {
	return db.Get(n, "select count(*) from duser")
}

// GetPortfolio returns the portfolio of a user
func (db *DB) GetPortfolio(portfolio *Portfolio, uuid int) error {
	return db.Get(portfolio, "select cash_bal, net_worth, rank, no_trans, margin from portfolio where user_id = $1", uuid)
}

// GetTickerData returns stock data for the ticker
func (db *DB) GetTickerData(tickerData *[]TickerData) error {
	return db.Select(tickerData, "select name, current_price, change_per from stocks_data")
}

// CreatePortfolio creates a portfolio for a new user
func (db *DB) CreatePortfolio(uuid int) (sql.Result, error) {
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

// GetStockHoldingsBuy gets the details about stocks bought by a certain user
func (db *DB) GetStockHoldingsBuy(userID int, stock *[]Stock) error {
	err := db.Select(stock, "select tr.symbol, tr.quantity, tr.value, s.current_price from transaction_buy as tr, stocks_data as s where tr.user_id = $1 and tr.symbol=s.symbol", userID)
	if err != nil {
		return err
	}
	for _, v := range *stock {
		v.Type = "BUY"
	}
	return nil
}

// GetStockHoldingsShortSell gets the details about stocks shorted by a certain user
func (db *DB) GetStockHoldingsShortSell(userID int, stock *[]Stock) error {
	err := db.Select(stock, "select tr.symbol, tr.quantity, tr.value, s.current_price from transaction_short_sell as tr, stocks_data as s where tr.user_id = $1 and tr.symbol=s.symbol", userID)
	if err != nil {
		return err
	}
	for _, v := range *stock {
		v.Type = "SHORT SELL"
	}
	return nil
}

// GetLeaderboard gets the users list in the descending order of level,
// and for users on the same level, in the ascending order of last submission time.
func (db *DB) GetLeaderboard(users *[]User) error {
	return db.Select(users, "select name, curr_level from duser order by curr_level desc, last_anstime")
}

// GetCompanyStockInfo returns all the details about the stock of a certain company
func (db *DB) GetCompanyStockInfo(companySymbol string, companyInfo *CompanyInfo) error {
	return db.Get(companyInfo, "select * from stocks_data where symbol = $1", companySymbol)
}

// GetPendingStocks returns info about pending stocks
func (db *DB) GetPendingStocks(uid int, pending *[]PendingData) error {
	return db.Select(pending, "select p.quantity, p.symbol, p.buy_ss, p.value, p.id, s.current_price from pending as p left join stocks_data as s on p.symbol = s.symbol where p.uid = $1", uid)
}

// DeletePending removes the specified entry from the pending table
func (db *DB) DeletePending(uid int, symbol string) (sql.Result, error) {
	return db.Exec("delete from pending where uid = $1 and symbol = $2", uid, symbol)
}

// GetCompanySymbols gets the stock exchange listing symbol for all the companies
func (db *DB) GetCompanySymbols(symbols *[]string) error {
	return db.Select(symbols, "select symbol from stocks_data")
}

// GetCurrentPrice returns the current price of a company's stock
func (db *DB) GetCurrentPrice(company string, currPrice *float32) error {
	return db.Get(currPrice, "select current_price from stocks_data where symbol = $1", company)
}

// GetUserInfoForCurrentPrice gets some additional user info to be sent whenever a company's current price is requested
func (db *DB) GetUserInfoForCurrentPrice(uid int, curr *CurrentPriceInfo) error {
	return db.Get(curr, "select cash_bal, margin, no_trans from portfolio where user_id = $1", uid)
}

// GetHistory gets the history of transactions of a user
func (db *DB) GetHistory(uid int, histories *[]History) error {
	err := db.Select(histories, "select * from history where uid =  $1", uid)
	if err != nil {
		return err
	}
	for _, v := range *histories {
		v.Total = v.Quantity * v.Price
	}
	return nil
}

// GetStockDataHistory gets the stock history of a specific company
func (db *DB) GetStockDataHistory(symbol string, stockHistoryData *[]StockHistory) error {
	err := db.Select(stockHistoryData, "select current_price, time from stock_data_history where symbol = $1 order by time desc limit 20", symbol)
	if err != nil {
		return err
	}
	return nil
}

// GetTransactionBuy gets the transactions where type is buy
func (db *DB) GetTransactionBuy(userID int, company string, txn *TransactionBuy) error {
	err := db.Get(&txn, "select * from transaction_buy where user_id = $1 and symbol = $2", userID, company)
	if err != nil {
		return err
	}
	return nil
}

// CreateNewTransactionBuy creates a new transaction for buying
func (db *DB) CreateNewTransactionBuy(userID int, company string, qty int, value float32) error {
	_, err := db.Exec("insert into transaction_buy values($1, $2, $3, $4)", userID, company, qty, value)
	return err
}

// UpdateTransactionBuy updates a transaction for buying
func (db *DB) UpdateTransactionBuy(userID int, company string, qty int, value float32) error {
	_, err := db.Exec("update transaction_buy set quantity = $3, value = $4 where user_id = $1 and symbol = $2", userID, company, qty, value)
	return err
}

// GetTransactionShortSell gets the transactions where type is short sell
func (db *DB) GetTransactionShortSell(userID int, company string, txn *TransactionShortSell) error {
	err := db.Get(&txn, "select * from transaction_buy where user_id = $1 and symbol = $2", userID, company)
	if err != nil {
		return err
	}
	return nil
}

// CreateNewTransactionShortSell creates a new transaction for short selling
func (db *DB) CreateNewTransactionShortSell(userID int, company string, qty int, value float32) error {
	_, err := db.Exec("insert into transaction_buy values($1, $2, $3, $4)", userID, company, qty, value)
	return err
}

// UpdateTransactionShortSell updates a short sell transaction
func (db *DB) UpdateTransactionShortSell(userID int, company string, qty int, value float32) error {
	_, err := db.Exec("update transaction_buy set quantity = $3, value = $4 where user_id = $1 and symbol = $2", userID, company, qty, value)
	return err
}

// UpdatePortfolio updates portfolio
func (db *DB) UpdatePortfolio(uuid int, cashBal, noOfTxns, margin float32) error {
	_, err := db.Exec("update portfolio set cash_bal = $2,  no_trans = $3, margin = $4 where user_id = $1", uuid, cashBal, noOfTxns, margin)
	return err
}

// CreateNewHistory creates a new history instance
func (db *DB) CreateNewHistory(userID int, h History) error {
	_, err := db.Exec("insert into history values($1, $2, $3, $4)", userID, h.Symbol, h.BuySS, h.Quantity, h.Price)
	return err
}
