package repository

import "it-league-stats/domain/model"

type PlayerRepository interface {
	SetupPlayers() ([]model.Player, error)
}
