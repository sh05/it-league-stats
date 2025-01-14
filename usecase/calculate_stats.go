package usecase

import (
	"it-league-stats/domain/model"
	"it-league-stats/domain/repository"
	"it-league-stats/domain/service"
)

func CalculateStats(gameRepo repository.GameRepository, playerRepo repository.PlayerRepository) (map[model.PlayerID]model.BattingResult, map[model.PlayerID]model.PitchingResult, error) {
	games, err := gameRepo.SetupGames()
	if err != nil {
		return nil, nil, err
	}

	players, err := playerRepo.SetupPlayers()
	if err != nil {
		return nil, nil, err
	}

	for _, player := range players {
		player.UpdateResults(games)
	}

	calculator := service.NewStatsCalculator()
	battingStats, pitchingStats := calculator.CalculateSeasonStats(games)

	return battingStats, pitchingStats, nil
}
