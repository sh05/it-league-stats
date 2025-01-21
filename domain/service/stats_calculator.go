package service

import (
	"it-league-stats/domain/model"
)

type StatsCalculator struct {
	// 開催試合数
	gamesHeld int
	// チーム打率 = 1試合でも出場した選手の打率の和 / 1試合でも出場した選手数
	leisuresAvg float64
}

func NewStatsCalculator(games []model.Game) *StatsCalculator {
	// playersの打率の平均を計算
	leisuresAvg := 0.0
	leisuresAtBats := 0.0
	leisuresHits := 0.0

	for _, game := range games {
		for _, br := range game.BattingResults {
			leisuresAtBats = leisuresAtBats + float64(br.PlateAppearances-br.Walks-br.HitsByPitch-br.SacrificeBunts-br.SacrificeFlies)
			leisuresHits = leisuresHits + float64(br.Singles+br.Doubles+br.Triples+br.HomeRuns)
		}
	}
	leisuresAvg = leisuresHits / leisuresAtBats
	return &StatsCalculator{
		gamesHeld:   len(games),
		leisuresAvg: leisuresAvg,
	}
}

func (sc *StatsCalculator) LeisuresAvg() float64 {
	return sc.leisuresAvg
}

func (sc *StatsCalculator) GamesHeld() int {
	return sc.gamesHeld
}

// 打数 = 打席数 - 四球 - 死球 - 犠打 - 犠飛
func (sc *StatsCalculator) CalculateAtBat(p model.Player) float64 {
	br := p.BattingResults
	return br.PlateAppearances - br.Walks - br.HitsByPitch - br.SacrificeBunts - br.SacrificeFlies
}

// 安打数 = 単打 + 二塁打 + 三塁打 + 本塁打
func (sc *StatsCalculator) CalculateHits(p model.Player) float64 {
	br := p.BattingResults
	return br.Singles + br.Doubles + br.Triples + br.HomeRuns
}

// 出塁率 = (安打数 + 四球 + 死球) / (打数 - 犠打 - 犠飛)
func (sc *StatsCalculator) CalculateOnBasePercentage(p model.Player) float64 {
	br := p.BattingResults
	return float64(sc.CalculateHits(p)+br.Walks+br.HitsByPitch) / float64(sc.CalculateAtBat(p))
}

// 打率 = 安打数 / 打数
func (sc *StatsCalculator) CalculateAverage(p model.Player) float64 {
	return float64(sc.CalculateHits(p)) / float64(sc.CalculateAtBat(p))
}

// POP STAR(フライ王) = 内野フライ + 外野フライ
func (sc *StatsCalculator) CalculatePopStar(p model.Player) float64 {
	br := p.BattingResults
	return br.InfieldFlies + br.OutfieldFlies
}

// No Power(外野まで飛ばない) = 内野ゴロ + 内野フライ
func (sc *StatsCalculator) CalculateNoPower(p model.Player) float64 {
	br := p.BattingResults
	return br.InfieldGrounders + br.InfieldFlies
}

// Mr. レジャーズ(チーム打率に最も近い打率)
// diff = チーム打率 - 打率
func (sc *StatsCalculator) CalculateDiffForMrLeisures(p model.Player) float64 {
	// チーム打率と打率の差の絶対値を取る
	diff := sc.leisuresAvg - sc.CalculateAverage(p)
	if diff < 0 {
		diff = -diff
	}
	return diff
}

// 凡打王 = 内野ゴロ + 内野フライ + 外野フライ
func (sc *StatsCalculator) CalculateTotalEasyOuts(p model.Player) float64 {
	br := p.BattingResults
	return br.InfieldGrounders + br.InfieldFlies + br.OutfieldFlies
}

// LOps = 出場率 + 出塁率
func (sc *StatsCalculator) CalculateLOps(p model.Player) float64 {
	// 出場率 = 出場試合数 / 全試合数
	playRate := float64(p.GamesPlayed) / float64(sc.gamesHeld)
	return playRate + sc.CalculateOnBasePercentage(p)
}

// BattingResultの各フィールドのget関数を実装
func (sc *StatsCalculator) Singles(p model.Player) float64 {
	return p.BattingResults.Singles
}

func (sc *StatsCalculator) Doubles(p model.Player) float64 {
	return p.BattingResults.Doubles
}

func (sc *StatsCalculator) Triples(p model.Player) float64 {
	return p.BattingResults.Triples
}

func (sc *StatsCalculator) HomeRuns(p model.Player) float64 {
	return p.BattingResults.HomeRuns
}

func (sc *StatsCalculator) Walks(p model.Player) float64 {
	return p.BattingResults.Walks
}

func (sc *StatsCalculator) HitsByPitch(p model.Player) float64 {
	return p.BattingResults.HitsByPitch
}

func (sc *StatsCalculator) SacrificeBunts(p model.Player) float64 {
	return p.BattingResults.SacrificeBunts
}

func (sc *StatsCalculator) SacrificeFlies(p model.Player) float64 {
	return p.BattingResults.SacrificeFlies
}

func (sc *StatsCalculator) PlateAppearances(p model.Player) float64 {
	return p.BattingResults.PlateAppearances
}

func (sc *StatsCalculator) InfieldGrounders(p model.Player) float64 {
	return p.BattingResults.InfieldGrounders
}

func (sc *StatsCalculator) InfieldFlies(p model.Player) float64 {
	return p.BattingResults.InfieldFlies
}

func (sc *StatsCalculator) OutfieldFlies(p model.Player) float64 {
	return p.BattingResults.OutfieldFlies
}

func (sc *StatsCalculator) RunsBattedIn(p model.Player) float64 {
	return p.BattingResults.RunsBattedIn
}

func (sc *StatsCalculator) StolenBases(p model.Player) float64 {
	return p.BattingResults.StolenBases
}

func (sc *StatsCalculator) BatterStrikeouts(p model.Player) float64 {
	return p.BattingResults.Strikeouts
}

func (sc *StatsCalculator) GamesPlayed(p model.Player) float64 {
	return p.GamesPlayed
}

func (sc *StatsCalculator) InningsPitched(p model.Player) float64 {
	return p.PitchingResults.InningsPitched
}

func (sc *StatsCalculator) PitcherStrikeouts(p model.Player) float64 {
	return p.PitchingResults.Strikeouts
}

func (sc *StatsCalculator) RunsAllowed(p model.Player) float64 {
	return p.PitchingResults.RunsAllowed
}

func (sc *StatsCalculator) Wins(p model.Player) float64 {
	return float64(p.PitchingResults.Wins)
}

func (sc *StatsCalculator) Losses(p model.Player) float64 {
	return float64(p.PitchingResults.Losses)
}
