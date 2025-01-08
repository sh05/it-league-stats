package model

type Game struct {
	Date            string
	Stadium         string
	OpponentTeam    string
	ScoreBoard      ScoreBoard
	BattingResults  map[PlayerID]BattingResults
	BattingOrder    Order
	PitchingResults map[PlayerID]PitchingResults
	PitchingOrder   Order
	MVPs            []PlayerID
}

type Order []PlayerID
