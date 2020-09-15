package utils

// CalculateBrokerage calculates the brokerage based on the
// following brokerage rates :
// For no. of transactions <= 100, rate = 0.1%
// For no. of transactions <= 1000, rate = 0.15%
// Else, rate = 0.3%
func CalculateBrokerage(noOfTrans float32, quantity int, currPrice float32) float32 {
	var brokerage float32
	if noOfTrans <= 100 {
		brokerage = ((0.1 * 100) * currPrice) * float32(quantity)
	} else if noOfTrans <= 1000 {
		brokerage = ((0.15 * 100) * currPrice) * float32(quantity)
	} else {
		brokerage = ((0.3 * 100) * currPrice) * float32(quantity)
	}
	return brokerage
}