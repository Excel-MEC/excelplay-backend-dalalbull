package tasks

import (
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/utils"
	"github.com/sirupsen/logrus"
)

func broadcastTickerData() {
	if utils.IsMarketOpen() {
		// TODO: Update ticker data and push via websockets
	} else {
		logrus.Info("Market is closed - not broadcasting ticker data")
	}
}
