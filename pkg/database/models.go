package database

import (
	"time"
)

// User holds the details of a particular user
type User struct {
	Name string `json:"name" db:"name"`
}

// Portfolio struct is used to hold details of a user's portfolio
type Portfolio struct {
	CashBal  float32 `json:"cash_bal" db:"cash_bal"`
	NetWorth float32 `json:"net_worth" db:"net_worth"`
	Rank     int     `json:"rank" db:"rank"`
	NoTrans  float32 `json:"no_trans" db:"no_trans"`
	Margin   float32 `json:"margin" db:"margin"`
}

// TickerData struct is used to hold details of a stock to be shown in a ticker
type TickerData struct {
	Name         string  `json:"name" db:"name"`
	CurrentPrice float32 `json:"current_price" db:"current_price"`
	ChangePer    float32 `json:"change_per" db:"change_per"`
}

// Stock struct holds information about a particular stock
type Stock struct {
	Company  string  `json:"company" db:"symbol"`
	Number   int     `json:"number" db:"quantity"`
	Type     string  `json:"type"`
	Purchase float32 `json:"purchase" db:"value"`
	Current  float32 `json:"current" db:"current_price"`
}

// CompanyInfo struct holds all the information about stock of a company.
type CompanyInfo struct {
	Symbol     string  `json:"symbol" db:"symbol"`
	Name       string  `json:"name" db:"name"`
	CurrPrice  float32 `json:"current_price" db:"current_price"`
	High       float32 `json:"high" db:"high"`
	Low        float32 `json:"low" db:"low"`
	OpenPrice  float32 `json:"open_price" db:"open_price"`
	Change     float32 `json:"change" db:"change"`
	ChangePer  float32 `json:"change_per" db:"change_per"`
	TradeQty   float32 `json:"trade_qty" db:"trade_qty"`
	TradeValue float32 `json:"trade_value" db:"trade_value"`
}

// StockChange struct holds change information about a particular stock
type StockChange struct {
	Name      string  `json:"name" db:"symbol"`
	ChangePer float32 `json:"change_per" db:"change_per"`
}

// StockVolume struct holds information about trade quantity of a particular stock
type StockVolume struct {
	Name     string  `json:"name" db:"symbol"`
	TradeQty float32 `json:"trade_qty" db:"trade_qty"`
}

// StockValue struct holds information about trade value of a particular stock
type StockValue struct {
	Name       string  `json:"name" db:"symbol"`
	TradeValue float32 `json:"trade_value" db:"trade_value"`
}

// DashboardData struct is used to hold all the data sent to be displayed on a user's dashboard
type DashboardData struct {
	TotalUsers    int           `json:"total_users"`
	StockHoldings []Stock       `json:"stockholdings"`
	TopGainers    []StockChange `json:"topGainers"`
	TopLosers     []StockChange `json:"topLosers"`
	MostActiveVol []StockVolume `json:"mostActiveVol"`
	MostActiveVal []StockValue  `json:"mostActiveVal"`
}

// SellDataWrapper is a wrapper around SellData providing additional information
type SellDataWrapper struct {
	CClose  bool       `json:"cclose"`
	NoStock bool       `json:"no_stock"`
	Data    []SellData `json:"trans"`
}

// SellData holds information about stocks a user has
type SellData struct {
	Company      string  `json:"company"`
	TypeOfTrade  string  `json:"buy_ss"`
	ShareInHand  int     `json:"share_in_hand"`
	CurrentPrice float32 `json:"current_price"`
	Gain         float32 `json:"gain"`
	TypeOfTrans  string  `json:"type_of_trans"`
}

// PendingData holds information about pending stocks of a user
type PendingData struct {
	Symbol       string  `json:"symbol" db:"symbol"`
	Quantity     int     `json:"quantity" db:"quantity"`
	Type         string  `json:"type" db:"buy_ss"`
	Value        float32 `json:"value" db:"value"`
	CurrentPrice float32 `json:"current_price" db:"current_price"`
	ID           string  `json:"id" db:"uid"`
}

// CurrentPriceInfo holds information to be sent back when user selects a company
type CurrentPriceInfo struct {
	CurrPrce float32 `json:"curr_price"`
	CashBal  float32 `json:"cash_bal" db:"cash_bal"`
	Margin   float32 `json:"margin" db:"margin"`
	NoTrans  float32 `json:"no_trans" db:"no_trans"`
}

// History holds data about past transactions of a user
type History struct {
	Symbol   string    `json:"symbol" db:"symbol"`
	BuySS    string    `json:"buy_ss" db:"buy_ss"`
	Quantity float32   `json:"quantity" db:"quantity"`
	Price    float32   `json:"price" db:"price"`
	Time     time.Time `json:"time" db:"time"`
	Total    float32   `json:"total"`
}

// HistoryArr holds an array of History instances
type HistoryArr struct {
	Histories []History `json:"history"`
}

type TransactionBuy struct {
	Symbol   string    `json:"symbol" db:"symbol"`
	Quantity int       `json:"quantity" db:"quantity"`
	Value    float32   `json:"value" db:"value"`
	ID       string    `json:"id" db:"user_id"`
	Time     time.Time `json:"time" db:"time"`
}
