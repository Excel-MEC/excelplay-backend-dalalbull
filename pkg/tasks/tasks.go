package tasks

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

// InitTasks starts various repeating tasks
func InitTasks() {
	c := cron.New(cron.WithSeconds()) // Uses the quartz cron standard over the Linux one to allow scheduling by seconds
	// Repeat every 10 seconds. This works only if the number of seconds divides 60.
	// 10 divides 60, so this works fine.
	c.AddFunc("*/10 * * * * *", test)
	c.AddFunc("*/10 * * * * *", stockUpdate)
	c.AddFunc("*/10 * * * * *", networthUpdate)
	c.AddFunc("*/10 * * * * *", broadcastGraphData)
	c.AddFunc("*/10 * * * * *", broadcastTickerData)
	c.AddFunc("*/10 * * * * *", broadcastPortfolioData)

	c.Start()
}

func test() {
	fmt.Println("Cron works!")
}
