package main

import (
	"flag"
	"fmt"
	"it-league-stats/infrastructure/repository"
	"it-league-stats/usecase"
	"log"
	"os"

	godotenv "github.com/joho/godotenv"
)

const (
	DEFAULT_OWN_TEAM = "HOGE"
)

func main() {
	excelFile := flag.String("file", "", "Path to the Excel file")
	flag.Parse()

	if *excelFile == "" {
		log.Fatal("Excel file path is required")
	}

	ownTeam := DEFAULT_OWN_TEAM
	err := godotenv.Load()
	if err == nil {
		log.Println("Error loading .env file")
	} else {
		ownTeam = os.Getenv("OWN_TEAM")
	}

	gameRepo := repository.NewExcelGameRepository(*excelFile, ownTeam)
	battingStats, pitchingStats, err := usecase.CalculateStats(gameRepo)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Batting Stats:")
	for playerID, stats := range battingStats {
		fmt.Printf("Player %#v: %#v", playerID, stats)
	}

	fmt.Println("\nPitching Stats:")
	for playerID, stats := range pitchingStats {
		fmt.Printf("Player %#v: %#v", playerID, stats)
	}
}
