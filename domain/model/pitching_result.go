package model

import (
	"log"
	"reflect"
)

type WinLossRecord float64

type PitchingResult struct {
	InningsPitched float64
	Strikeouts     float64
	RunsAllowed    float64
	Wins           WinLossRecord
	Losses         WinLossRecord
}

func NewPitchingResult() PitchingResult {
	return PitchingResult{
		InningsPitched: 0,
		Strikeouts:     0,
		RunsAllowed:    0,
		Wins:           WinLossRecord(0),
		Losses:         WinLossRecord(0),
	}
}

func (pr *PitchingResult) Update(newResult PitchingResult) {
	prValue := reflect.ValueOf(pr).Elem()
	newValue := reflect.ValueOf(newResult)

	for i := 0; i < prValue.NumField(); i++ {
		field := prValue.Field(i)
		newField := newValue.Field(i)

		switch field.Kind() {
		case reflect.Float64:
			field.SetFloat(field.Float() + newField.Float())
		case reflect.Int:
			field.SetInt(field.Int() + newField.Int())
		default:
			log.Fatal("Unexpected Type")
		}
	}
}
