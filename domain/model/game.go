package model

type Game struct {
	Date            string
	Stadium         string
	OpponentTeam    string
	Score           Score
	BattingResults  map[PlayerID]BattingResults
	PitchingResults map[PlayerID]PitchingResults
	MVPs            []PlayerID
}
