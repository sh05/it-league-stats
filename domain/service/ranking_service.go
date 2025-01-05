package service

import (
	"sort"
	"sync"
)

type RankingService struct {
	items map[string]float64
	mu    sync.Mutex
}

func NewRankingService() *RankingService {
	return &RankingService{
		items: make(map[string]float64),
	}
}

func (rs *RankingService) AddScore(id string, score float64) {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	rs.items[id] = score
}

func (rs *RankingService) GetTopN(n int) []string {
	rs.mu.Lock()
	defer rs.mu.Unlock()

	type item struct {
		id    string
		score float64
	}

	items := make([]item, 0, len(rs.items))
	for id, score := range rs.items {
		items = append(items, item{id: id, score: score})
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].score > items[j].score
	})

	result := make([]string, 0, n)
	for i := 0; i < n && i < len(items); i++ {
		result = append(result, items[i].id)
	}

	return result
}
