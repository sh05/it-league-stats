package model

import (
	"log"
	"reflect"
)

type BattingResult struct {
	InfieldFlies     float64
	InfieldGrounders float64
	OutfieldFlies    float64
	Strikeouts       float64
	Walks            float64
	HitsByPitch      float64
	SacrificeBunts   float64
	SacrificeFlies   float64
	Singles          float64
	Doubles          float64
	Triples          float64
	HomeRuns         float64
	RunsBattedIn     float64
	StolenBases      float64
	PlateAppearances float64
}

func NewBattingResult() BattingResult {
	return BattingResult{
		InfieldFlies:     0,
		InfieldGrounders: 0,
		OutfieldFlies:    0,
		Strikeouts:       0,
		Walks:            0,
		HitsByPitch:      0,
		SacrificeBunts:   0,
		SacrificeFlies:   0,
		Singles:          0,
		Doubles:          0,
		Triples:          0,
		HomeRuns:         0,
		RunsBattedIn:     0,
		StolenBases:      0,
		PlateAppearances: 0,
	}
}

func (br *BattingResult) Update(newResult BattingResult) {
	brValue := reflect.ValueOf(br).Elem()
	newValue := reflect.ValueOf(newResult)

	for i := 0; i < brValue.NumField(); i++ {
		field := brValue.Field(i)
		newField := newValue.Field(i)

		switch field.Kind() {
		case reflect.Float64:
			field.SetFloat(field.Float() + newField.Float())
		default:
			log.Fatal("Unexpected Type")
		}
	}
}
