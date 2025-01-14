package model

type Game struct {
	Date            string
	Stadium         string
	OpponentTeam    string
	ScoreBoard      ScoreBoard
	BattingResults  map[PlayerID]BattingResult
	BattingOrder    Order
	PitchingResults map[PlayerID]PitchingResult
	PitchingOrder   Order
	MVPs            []PlayerID
}

type Order []PlayerID
