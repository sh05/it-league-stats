package model

type RankingEntry struct {
	PlayerID PlayerID
	Value    float64
}

type Ranking struct {
	Category string
	Entries  []RankingEntry
}

// N位までを返す
func (r *Ranking) topk(n int) []RankingEntry {
	return r.Entries[:n]
}
