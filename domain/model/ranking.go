package model

type RankingEntry struct {
	PlayerID PlayerID
	Value    float64
}

type Ranking struct {
	Category string
	Entries  []RankingEntry
}
