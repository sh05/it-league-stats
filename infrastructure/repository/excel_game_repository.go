package repository

import (
	"it-league-stats/domain/model"
	"it-league-stats/domain/repository"
	"it-league-stats/infrastructure/excel"
	"log"
	"regexp"
	"runtime"
	"slices"
	"strconv"
)

const (
	EXAMPLE_GAME_SHEET                          = "mmdd_対戦相手名"
	GAME_SHEET_PATTERN                          = `^\d{4}_.*$`
	DATE_ROW                                    = 2
	DATE_COL                                    = 2
	STADIUM_ROW                                 = 2
	STADIUM_COL                                 = 3
	SCORE_BOARD_ROW_FROM                        = 3
	SCORE_BOARD_ROW_TO                          = 4 + 1
	SCORE_BOARD_COL_FROM                        = 8
	SCORE_BOARD_COL_TO                          = 21 + 1
	SCORE_BOARD_TOP_ROW                         = 0
	SCORE_BOARD_BOTTOM_ROW                      = 1
	SCORE_BOARD_BATFIRST_COL                    = 0
	SCORE_BOARD_FIELDFIRST_COL                  = 0
	SCORE_BOARD_TOP_COL_FROM                    = 1
	SCORE_BOARD_TOP_COL_TO                      = 9
	SCORE_BOARD_BOTTOM_COL_FROM                 = 1
	SCORE_BOARD_BOTTOM_COL_TO                   = 9 + 1
	BATTING_RESULTS_ROW_FROM                    = 8
	BATTING_RESULTS_ROW_TO                      = 33 + 1
	BATTING_RESULTS_PLAYER_ID_COL               = 7
	BATTING_RESULTS_INFILD_FLY_COL              = 8
	BATTING_RESULTS_INFILD_GROUNDER_COL         = 9
	BATTING_RESULTS_OUTFIELD_FLY_COL            = 10
	BATTING_RESULTS_STRIKEOUT_COL               = 11
	BATTING_RESULTS_WALK_COL                    = 12
	BATTING_RESULTS_HIT_BY_PITCH_COL            = 13
	BATTING_RESULTS_SACRIFICE_BUNT_COL          = 14
	BATTING_RESULTS_SACRIFICE_FLY_COL           = 15
	BATTING_RESULTS_SINGLE_COL                  = 16
	BATTING_RESULTS_DOUBLE_COL                  = 17
	BATTING_RESULTS_TRIPLE_COL                  = 18
	BATTING_RESULTS_HOME_RUN_COL                = 19
	BATTING_RESULTS_RUNS_BATTED_IN_COL          = 20
	BATTING_RESULTS_STOLEN_BASE_COL             = 21
	BATTING_RESULTS_PLATE_APPEARANCES_COL       = 22
	PITCHING_RESULTS_ROW_FROM                   = 37
	PITCHING_RESULTS_ROW_TO                     = 42 + 1
	PITCHING_RESULTS_PLAYER_ID_COL              = 3
	PITCHING_RESULTS_WIN_COL                    = 4
	PITCHING_RESULTS_PITCHED_INNINGS_COL        = 5
	PITCHING_RESULTS_PITCHED_INNINGS_THIRDS_COL = 7
	PITCHING_RESULTS_STRIKEOUT_COL              = 9
	PITCHING_RESULTS_RUNS_ALLOWED_COL           = 11
	PITCHING_RESULTS_WIN_CHAR                   = "勝"
	PITCHING_RESULTS_LOSS_CHAR                  = "敗"
	MVP_ROW                                     = 44
	MVP_FIRST_PLAYER_COL                        = 3
	MVP_SECOND_PLAYER_COL                       = 10
)

type ExcelGameRepository struct {
	baseGameRepository repository.BaseGameRepository
	filePath           string
}

func NewExcelGameRepository(filePath, ownTeam string) *ExcelGameRepository {
	return &ExcelGameRepository{
		baseGameRepository: repository.BaseGameRepository{
			OwnTeam: ownTeam,
		},
		filePath: filePath,
	}
}

func (r *ExcelGameRepository) SetupGames() ([]model.Game, error) {
	games := []model.Game{}

	f, err := excel.ReadExcelFile(r.filePath)
	if err != nil {
		return nil, err
	}

	for _, sheetName := range f.GetSheetMap() {
		if !regexp.MustCompile(GAME_SHEET_PATTERN).MatchString(sheetName) {
			continue
		}

		sheet, err := excel.ReadExcelSheet(f, sheetName)
		if err != nil {
			return nil, err
		}
		games = append(games, parseGameSheet(sheetName, sheet))
	}
	return games, nil
}

func parseGameSheet(sheetName string, sheet [][]string) model.Game {
	date, opponentTeam := dateAndOpponent(sheetName)
	battingResults, battingOrder := parseBattingResultsTable(sheet[BATTING_RESULTS_ROW_FROM:BATTING_RESULTS_ROW_TO])
	pitchingResults, pitchingOrder := parsePitchingResultsTable(sheet[PITCHING_RESULTS_ROW_FROM:PITCHING_RESULTS_ROW_TO])
	return model.Game{
		Date:            date,
		Stadium:         sheet[STADIUM_COL][STADIUM_ROW],
		OpponentTeam:    opponentTeam,
		ScoreBoard:      parseScoreBoard(sheet[SCORE_BOARD_ROW_FROM:SCORE_BOARD_ROW_TO]),
		BattingResults:  battingResults,
		BattingOrder:    battingOrder,
		PitchingResults: pitchingResults,
		PitchingOrder:   pitchingOrder,
		MVPs:            parseMVPRow(sheet[MVP_ROW]),
	}
}

func parseScoreBoard(scoreBoardRow [][]string) model.ScoreBoard {
	offset := SCORE_BOARD_COL_FROM
	scoreBoardTopRow := scoreBoardRow[SCORE_BOARD_TOP_ROW]
	scoreBoardBottomRow := scoreBoardRow[SCORE_BOARD_BOTTOM_COL_FROM]

	return model.ScoreBoard{
		BatFirst: model.Score{
			Team:   scoreBoardTopRow[SCORE_BOARD_BATFIRST_COL+offset],
			Scores: strSlice2Float64Slice(scoreBoardTopRow[SCORE_BOARD_TOP_COL_FROM+offset : SCORE_BOARD_TOP_COL_TO+offset]),
		},
		FieldFirst: model.Score{
			Team:   scoreBoardBottomRow[SCORE_BOARD_FIELDFIRST_COL+offset],
			Scores: strSlice2Float64Slice(scoreBoardBottomRow[SCORE_BOARD_BOTTOM_COL_FROM+offset : SCORE_BOARD_BOTTOM_COL_TO+offset]),
		},
	}
}

func parseBattingResultsTable(table [][]string) (map[model.PlayerID]model.BattingResult, model.Order) {
	battingResults := make(map[model.PlayerID]model.BattingResult)
	battingOrder := model.Order{}
	for _, row := range table {
		playerID := model.PlayerID(row[BATTING_RESULTS_PLAYER_ID_COL])
		if playerID == "" {
			continue
		}
		battingOrder = append(battingOrder, playerID)
		battingResults[playerID] = model.BattingResult{
			InfieldFlies:     str2Float64Easily(row[BATTING_RESULTS_INFILD_FLY_COL]),
			InfieldGrounders: str2Float64Easily(row[BATTING_RESULTS_INFILD_GROUNDER_COL]),
			OutfieldFlies:    str2Float64Easily(row[BATTING_RESULTS_OUTFIELD_FLY_COL]),
			Strikeouts:       str2Float64Easily(row[BATTING_RESULTS_STRIKEOUT_COL]),
			Walks:            str2Float64Easily(row[BATTING_RESULTS_WALK_COL]),
			HitsByPitch:      str2Float64Easily(row[BATTING_RESULTS_HIT_BY_PITCH_COL]),
			SacrificeBunts:   str2Float64Easily(row[BATTING_RESULTS_SACRIFICE_BUNT_COL]),
			SacrificeFlies:   str2Float64Easily(row[BATTING_RESULTS_SACRIFICE_FLY_COL]),
			Singles:          str2Float64Easily(row[BATTING_RESULTS_SINGLE_COL]),
			Doubles:          str2Float64Easily(row[BATTING_RESULTS_DOUBLE_COL]),
			Triples:          str2Float64Easily(row[BATTING_RESULTS_TRIPLE_COL]),
			HomeRuns:         str2Float64Easily(row[BATTING_RESULTS_HOME_RUN_COL]),
			RunsBattedIn:     str2Float64Easily(row[BATTING_RESULTS_RUNS_BATTED_IN_COL]),
			StolenBases:      str2Float64Easily(row[BATTING_RESULTS_STOLEN_BASE_COL]),
			PlateAppearances: str2Float64Easily(row[BATTING_RESULTS_PLATE_APPEARANCES_COL]),
		}
	}
	return battingResults, battingOrder
}

func parsePitchingResultsTable(table [][]string) (map[model.PlayerID]model.PitchingResult, model.Order) {
	pitchingResults := make(map[model.PlayerID]model.PitchingResult)
	pitchingOrder := model.Order{}
	for _, row := range table {
		playerID := model.PlayerID(row[PITCHING_RESULTS_PLAYER_ID_COL])
		if playerID == "" {
			continue
		}

		pitchingOrder = append(pitchingOrder, playerID)
		wins := model.WinLossRecord(0)
		losses := model.WinLossRecord(0)
		switch row[PITCHING_RESULTS_WIN_COL] {
		case PITCHING_RESULTS_WIN_CHAR:
			wins = model.WinLossRecord(1)
			losses = model.WinLossRecord(0)
		case PITCHING_RESULTS_LOSS_CHAR:
			wins = model.WinLossRecord(0)
			losses = model.WinLossRecord(1)
		default:
			wins = model.WinLossRecord(0)
			losses = model.WinLossRecord(0)
		}

		runsAllowed := 0.0
		if len(row) >= PITCHING_RESULTS_RUNS_ALLOWED_COL {
			runsAllowed = str2Float64Easily(row[PITCHING_RESULTS_RUNS_ALLOWED_COL])
		}
		pitchingResults[playerID] = model.PitchingResult{
			Strikeouts:     str2Float64Easily(row[PITCHING_RESULTS_STRIKEOUT_COL]),
			InningsPitched: float64(str2Float64Easily(row[PITCHING_RESULTS_PITCHED_INNINGS_COL]) + str2Float64Easily(row[PITCHING_RESULTS_PITCHED_INNINGS_THIRDS_COL])/10),
			RunsAllowed:    runsAllowed,
			Wins:           wins,
			Losses:         losses,
		}

	}
	return pitchingResults, pitchingOrder
}

func parseMVPRow(row []string) []model.PlayerID {
	return []model.PlayerID{
		(model.PlayerID)(row[MVP_FIRST_PLAYER_COL]),
		(model.PlayerID)(row[MVP_SECOND_PLAYER_COL]),
	}
}

func strSlice2Float64Slice(s []string) []float64 {
	var floatSlice []float64
	for _, str := range s {
		floatSlice = append(floatSlice, str2Float64Easily(str))
	}
	return floatSlice
}

func str2Float64Easily(s string) float64 {
	skipChars := []string{"", "x"}
	var f float64
	if slices.Contains(skipChars, s) {
		s = "0"
	}

	// sをfloat64にする処理
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Print(s)
		log.Print(f)
		log.Panic(err)
		runtime.Goexit()
	}
	return f
}

func dateAndOpponent(sheetName string) (string, string) {
	return sheetName[:4], sheetName[5:]
}
