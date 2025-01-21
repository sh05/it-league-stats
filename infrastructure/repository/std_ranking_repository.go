package repository

import (
	"fmt"
	"it-league-stats/domain/model"
)

type StdoutRankingRepository struct{}

func NewStdoutRankingRepository() *StdoutRankingRepository {
	return &StdoutRankingRepository{}
}

func (r *StdoutRankingRepository) Write(ranking model.Ranking, players []model.Player) {
	header := []interface{}{fmt.Sprintf("%sの順位", ranking.Category), "選手", "値"}
	fmt.Println(header)
	for i, e := range ranking.Entries {
		player := model.PlayerByID(players, e.PlayerID)
		fmt.Println([]interface{}{i + 1, player.NameAndId(), e.Value})
	}
}
