package repository

import (
	"it-league-stats/domain/model"
)

const (
	PLAYER_LIST_SHEET = "選手名簿"
)

type ExcelPlayerRepository struct {
	filePath string
}

func NewExcelPlayerRepository(filePath string) *ExcelPlayerRepository {
	return &ExcelPlayerRepository{filePath: filePath}
}

func (r *ExcelPlayerRepository) GetAllPlayers() ([]model.Player, error) {
	// Excelファイルを読み込み、Playerオブジェクトのスライスを返す
	return nil, nil
}
