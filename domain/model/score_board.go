package model

type ScoreBoard struct {
	BatFirst   Score
	FieldFirst Score
}

type Score struct {
	Team   string
	Scores []int
}
