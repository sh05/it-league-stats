package repository

import "it-league-stats/domain/model"

type GameRepository interface {
	GetAllGames() ([]model.Game, error)
}
