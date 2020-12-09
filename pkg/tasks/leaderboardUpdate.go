package tasks

import (
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/utils"
	"github.com/sirupsen/logrus"
)

// TODO: This is apparently not necessary because the leaderboard is ordered by net worth
// Remove after testing this is true
func leaderboardUpdate() {
	if utils.IsMarketOpen() {
		logrus.Info("Updating leaderboard...")
	}
}
