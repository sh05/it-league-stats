package model

type WinLossRecord int

type PitchingResults struct {
	InningsPitched float64
	Strikeouts     int
	RunsAllowed    int
	Wins           WinLossRecord
	Losses         WinLossRecord
}
