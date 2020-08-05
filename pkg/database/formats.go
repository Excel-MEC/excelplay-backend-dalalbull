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
