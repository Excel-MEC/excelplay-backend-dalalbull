package tasks

import (
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/utils"
	"github.com/sirupsen/logrus"
)

func broadcastGraphData() {
	if utils.IsMarketOpen() {
		// TODO: Update graph and push via websockets
	} else {
		logrus.Info("Market is closed - not broadcasting graph updates")
	}
}
