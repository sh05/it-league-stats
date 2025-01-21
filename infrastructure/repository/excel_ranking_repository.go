package repository

import (
	"fmt"
	"it-league-stats/domain/model"
	"it-league-stats/infrastructure/excel"
)

type ExcelRankingRepository struct {
	filePath string
}

func NewExcelRankingRepository(filePath string) *ExcelRankingRepository {
	excel.TouchExcelFile(filePath)
	return &ExcelRankingRepository{filePath: filePath}
}

func (r *ExcelRankingRepository) Write(ranking model.Ranking, players []model.Player) {
	printData := [][]interface{}{}
	header := []interface{}{fmt.Sprintf("%sの順位", ranking.Category), "選手", "値"}
	printData = append(printData, header)
	for i, e := range ranking.Entries {
		player := model.PlayerByID(players, e.PlayerID)
		printData = append(printData, []interface{}{i + 1, player.NameAndId(), e.Value})
	}
	excel.WriteExcelSheet(r.filePath, ranking.Category, printData)
}
