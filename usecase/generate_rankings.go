package usecase

import (
	"fmt"
	"it-league-stats/domain/model"
	"it-league-stats/domain/repository"
	"it-league-stats/domain/service"
	"log"
	"sort"
)

type RankingGenerator struct {
	calculator  *service.StatsCalculator
	players     []model.Player
	games       []model.Game
	rankingRepo repository.RankingRepository
}

func NewRankingGenerator(playerRepo repository.PlayerRepository, gameRepo repository.GameRepository, rankingRepo repository.RankingRepository) (*RankingGenerator, error) {
	games, players, err := Setup(gameRepo, playerRepo)
	if err != nil {
		return nil, err
	}
	return &RankingGenerator{
		players:     players,
		games:       games,
		calculator:  service.NewStatsCalculator(games),
		rankingRepo: rankingRepo,
	}, nil
}

func (rg *RankingGenerator) PrintRankings() {
	rankings, err := rg.GenerateRankings()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		return
	}
	for _, r := range rankings {
		rg.rankingRepo.Write(r, rg.players)
	}
}

func (rg *RankingGenerator) GenerateRankings() ([]model.Ranking, error) {
	rankings := []model.Ranking{
		rg.generateRanking("首位打者", rg.calculator.CalculateAverage, true),
		rg.generateRanking("ホームラン王", rg.calculator.HomeRuns, true),
		rg.generateRanking("最多安打", rg.calculator.CalculateHits, true),
		rg.generateRanking(fmt.Sprintf("出場試合数(2024年度試合数 %d)", rg.calculator.GamesHeld()), rg.calculator.GamesPlayed, true),
		rg.generateRanking("ノーパワー", rg.calculator.CalculateNoPower, true),
		rg.generateRanking("POP Star", rg.calculator.CalculatePopStar, true),
		rg.generateRanking("最優秀凡打王", rg.calculator.CalculateTotalEasyOuts, true),
		rg.generateRanking(fmt.Sprintf("Mr.レジャーズ(チーム打率 %f)", rg.calculator.LeisuresAvg()), rg.calculator.CalculateDiffForMrLeisures, false),
		rg.generateRanking("LOps(出場率+出塁率)", rg.calculator.CalculateLOps, true),
		rg.generateRanking("三振王", rg.calculator.BatterStrikeouts, true),
		rg.generateRanking("打点王", rg.calculator.RunsBattedIn, true),
		rg.generateRanking("最多勝", rg.calculator.Wins, true),
		rg.generateRanking("奪三振王", rg.calculator.PitcherStrikeouts, true),
		rg.generateRanking("鉄腕(最多イニング投球)", rg.calculator.Wins, true),
		rg.generateRanking("お散歩キング", rg.calculator.Walks, true),
		rg.generateRanking("盗塁王", rg.calculator.StolenBases, true),
	}

	return rankings, nil
}

func (rg *RankingGenerator) generateRanking(category string, f func(model.Player) float64, inverse bool) model.Ranking {
	entries := make([]model.RankingEntry, 0, len(rg.players))
	for _, p := range rg.players {
		entries = append(entries, model.RankingEntry{PlayerID: p.ID, Value: f(p)})
	}
	sort.SliceStable(entries, func(i, j int) bool {
		if inverse {
			return entries[i].Value > entries[j].Value
		}
		return entries[i].Value < entries[j].Value
	})
	return model.Ranking{Category: category, Entries: entries}
}
