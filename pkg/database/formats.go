package database

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
