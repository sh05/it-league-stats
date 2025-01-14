package model

import (
	"fmt"
)

type PlayerID string

type Player struct {
	ID              PlayerID
	Name            string
	GamesPlayed     int
	BattingResults  BattingResult
	PitchingResults PitchingResult
}

func NewPlayer(id PlayerID, name string) Player {
	return Player{
		ID:              id,
		Name:            name,
		GamesPlayed:     0,
		BattingResults:  NewBattingResult(),
		PitchingResults: NewPitchingResult(),
	}
}

func (p *Player) NameAndId() string {
	return fmt.Sprintf("%s#%s", p.Name, p.ID)
}

func (p *Player) UpdateResults(games []Game) {
	for _, game := range games {
		participated := false
		if br, ok := game.BattingResults[p.ID]; ok {
			participated = true
			p.BattingResults.Update(br)
		}
		if pr, ok := game.PitchingResults[p.ID]; ok {
			participated = true
			p.PitchingResults.Update(pr)
		}
		if participated {
			p.GamesPlayed += 1
		}
	}
}
