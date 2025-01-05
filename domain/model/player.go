package model

type PlayerID string

type Player struct {
	Name            string
	ID              PlayerID
	GamesPlayed     int
	BattingResults  BattingResults
	PitchingResults PitchingResults
}
