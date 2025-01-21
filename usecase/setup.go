package usecase

import (
	"it-league-stats/domain/model"
	"it-league-stats/domain/repository"
)

func Setup(gameRepo repository.GameRepository, playerRepo repository.PlayerRepository) ([]model.Game, []model.Player, error) {
	games, err := gameRepo.SetupGames()
	if err != nil {
		return nil, nil, err
	}

	players, err := playerRepo.SetupPlayers()
	if err != nil {
		return nil, nil, err
	}

	updatedPlayers := make([]model.Player, 0, len(players))
	for _, player := range players {
		player.UpdateResults(games)
		if player.GamesPlayed > 0 {
			updatedPlayers = append(updatedPlayers, player)
		}
	}

	return games, updatedPlayers, nil
}
