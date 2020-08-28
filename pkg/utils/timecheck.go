package utils

import (
	"time"
)

// IsMarketOpen checks if the stock market is open at the current moment
func IsMarketOpen() bool {
	now := time.Now()
	// Stock market is not open on Sunday and Saturday
	if int(now.Weekday()) == 0 || int(now.Weekday()) == 6 {
		return false
	}
	// Check if stock market is open
	if (now.Hour() < 9 || (now.Hour() == 9 && now.Minute() < 15)) || (now.Hour() > 15 || (now.Hour() == 15 && now.Minute() > 30)) {
		return false
	}
	return true
}
