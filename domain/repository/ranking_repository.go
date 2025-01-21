package repository

import "it-league-stats/domain/model"

type RankingRepository interface {
	Write(model.Ranking, []model.Player)
}
