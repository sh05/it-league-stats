package usecase

import (
	"it-league-stats/domain/model"
	"it-league-stats/domain/repository"
	"it-league-stats/domain/service"
)

func CalculateStats(gameRepo repository.GameRepository) (map[model.PlayerID]model.BattingResults, map[model.PlayerID]model.PitchingResults, error) {
	games, err := gameRepo.GetAllGames()
	if err != nil {
		return nil, nil, err
	}

	calculator := service.NewStatsCalculator()
	battingStats, pitchingStats := calculator.CalculateSeasonStats(games)

	return battingStats, pitchingStats, nil
}

func AllData(gameRepo repository.GameRepository, playerRepo repository.PlayerRepository) ([]model.Game, []model.Player, error) {
	games, err := gameRepo.GetAllGames()
	if err != nil {
		return nil, nil, err
	}

	players, err := playerRepo.GetAllPlayers()
	if err != nil {
		return nil, nil, err
	}

	return games, players, nil
}
