package tasks

import (
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/utils"
	"github.com/sirupsen/logrus"
)

func broadcastPortfolioData() {
	if utils.IsMarketOpen() {
		// TODO: Update portfolio data and push via websockets
	} else {
		logrus.Info("Market is closed - not broadcasting portfolio data")
	}
}
