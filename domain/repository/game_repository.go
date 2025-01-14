package repository

import "it-league-stats/domain/model"

type GameRepository interface {
	SetupGames() ([]model.Game, error)
}

type BaseGameRepository struct {
	OwnTeam string
}
