package repository

import "it-league-stats/domain/model"

type PlayerRepository interface {
	GetAllPlayers() ([]model.Player, error)
}
