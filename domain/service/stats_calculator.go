package service

import (
	"it-league-stats/domain/model"
)

type StatsCalculator struct{}

func NewStatsCalculator() *StatsCalculator {
	return &StatsCalculator{}
}

func (sc *StatsCalculator) CalculateSeasonStats(games []model.Game) (map[model.PlayerID]model.BattingResults, map[model.PlayerID]model.PitchingResults) {
	battingStats := make(map[model.PlayerID]model.BattingResults)
	pitchingStats := make(map[model.PlayerID]model.PitchingResults)

	for _, game := range games {
		sc.updateBattingStats(battingStats, game)
		sc.updatePitchingStats(pitchingStats, game)
	}

	sc.calculateAverages(battingStats)
	sc.calculateERAs(pitchingStats)

	return battingStats, pitchingStats
}

func (sc *StatsCalculator) updateBattingStats(stats map[model.PlayerID]model.BattingResults, game model.Game) {
	for playerID, gameStat := range game.BattingResults {
		playerStats, exists := stats[playerID]
		if !exists {
			playerStats = model.BattingResults{}
		}

		playerStats.Singles += gameStat.Singles
		playerStats.RunsBattedIn += gameStat.RunsBattedIn
		playerStats.HomeRuns += gameStat.HomeRuns
		playerStats.Walks += gameStat.Walks

		stats[playerID] = playerStats
	}
}

func (sc *StatsCalculator) updatePitchingStats(stats map[model.PlayerID]model.PitchingResults, game model.Game) {
	for playerID, gameStat := range game.PitchingResults {
		playerStats, exists := stats[playerID]
		if !exists {
			playerStats = model.PitchingResults{}
		}

		playerStats.InningsPitched += gameStat.InningsPitched
		playerStats.RunsAllowed += gameStat.RunsAllowed
		playerStats.Strikeouts += gameStat.Strikeouts

		// if gameStat.Win == model.Win {
		// 	playerStats.Wins++
		// } else if gameStat.WinLoss == model.Loss {
		// 	playerStats.Losses++
		// }

		stats[playerID] = playerStats
	}
}

func (sc *StatsCalculator) calculateAverages(stats map[model.PlayerID]model.BattingResults) {
	for playerID, stat := range stats {
		// if stat.AtBats > 0 {
		// 	stat.BattingAverage = float64(stat.Hits) / float64(stat.AtBats)
		// }
		stats[playerID] = stat
	}
}

func (sc *StatsCalculator) calculateERAs(stats map[model.PlayerID]model.PitchingResults) {
	for playerID, stat := range stats {
		// if stat.InningsPitched > 0 {
		// 	stat.ERA = (stat.EarnedRuns * 9) / stat.InningsPitched
		// }
		stats[playerID] = stat
	}
}
