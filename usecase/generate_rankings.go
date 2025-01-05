package usecase

import (
	"it-league-stats/domain/model"
	"it-league-stats/domain/repository"
	"it-league-stats/domain/service"
	"sort"
)

type RankingGenerator struct {
	playerRepo repository.PlayerRepository
	gameRepo   repository.GameRepository
	calculator *service.StatsCalculator
}

func NewRankingGenerator(playerRepo repository.PlayerRepository, gameRepo repository.GameRepository) *RankingGenerator {
	return &RankingGenerator{
		playerRepo: playerRepo,
		gameRepo:   gameRepo,
		calculator: service.NewStatsCalculator(),
	}
}

func (rg *RankingGenerator) GenerateRankings() ([]model.Ranking, error) {
	games, err := rg.gameRepo.GetAllGames()
	if err != nil {
		return nil, err
	}

	battingStats, pitchingStats := rg.calculator.CalculateSeasonStats(games)

	rankings := []model.Ranking{
		rg.generateHomeRunRanking(battingStats),
		rg.generateStrikeoutRanking(pitchingStats),
	}

	return rankings, nil
}

func (rg *RankingGenerator) generateHomeRunRanking(stats map[model.PlayerID]model.BattingResults) model.Ranking {
	entries := make([]model.RankingEntry, 0, len(stats))
	for playerID, stat := range stats {
		entries = append(entries, model.RankingEntry{PlayerID: playerID, Value: float64(stat.HomeRuns)})
	}
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Value > entries[j].Value
	})
	return model.Ranking{Category: "本塁打", Entries: entries[:min(len(entries), 10)]}
}

func (rg *RankingGenerator) generateStrikeoutRanking(stats map[model.PlayerID]model.PitchingResults) model.Ranking {
	entries := make([]model.RankingEntry, 0, len(stats))
	for playerID, stat := range stats {
		entries = append(entries, model.RankingEntry{PlayerID: playerID, Value: float64(stat.Strikeouts)})
	}
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Value > entries[j].Value
	})
	return model.Ranking{Category: "奪三振", Entries: entries[:min(len(entries), 10)]}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
