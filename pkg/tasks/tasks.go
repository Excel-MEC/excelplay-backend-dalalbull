package tasks

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

// InitTasks starts various repeating tasks
func InitTasks() {
	c := cron.New(cron.WithSeconds())
	// Repeat every 10 seconds. This works only if the number of seconds divides 60.
	// 10 divides 60, so this works fine.
	c.AddFunc("*/10 * * * * *", test)

	c.Start()
}

func test() {
	fmt.Println("Cron works!")
}
