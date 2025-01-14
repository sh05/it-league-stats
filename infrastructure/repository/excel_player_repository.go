package repository

import (
	"it-league-stats/domain/model"
	"it-league-stats/infrastructure/excel"
)

const (
	PLAYER_LIST_SHEET    = "選手名簿"
	PLAYER_LIST_ROW_FROM = 0
	PLAYER_LIST_ROW_TO   = 102
	PLAYER_LIST_COL_FROM = 0
	PLAYER_LIST_COL_TO   = 1
)

type ExcelPlayerRepository struct {
	filePath string
}

func NewExcelPlayerRepository(filePath string) *ExcelPlayerRepository {
	return &ExcelPlayerRepository{filePath: filePath}
}

func (r *ExcelPlayerRepository) SetupPlayers() ([]model.Player, error) {
	f, err := excel.ReadExcelFile(r.filePath)
	if err != nil {
		return nil, err
	}

	sheet, err := excel.ReadExcelSheet(f, PLAYER_LIST_SHEET)
	if err != nil {
		return nil, err
	}

	return parsePlayerList(sheet), nil
}

func parsePlayerList(sheet [][]string) []model.Player {
	players := []model.Player{}
	for _, row := range sheet[PLAYER_LIST_ROW_FROM:PLAYER_LIST_ROW_TO] {
		if len(row) < 2 {
			continue
		}
		players = append(players, model.NewPlayer(model.PlayerID(row[PLAYER_LIST_COL_FROM]), row[PLAYER_LIST_COL_TO]))
	}
	return players
}
