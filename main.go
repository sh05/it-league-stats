package main

import (
	"flag"
	drepository "it-league-stats/domain/repository"
	irepository "it-league-stats/infrastructure/repository"
	"it-league-stats/usecase"
	"log"
	"os"

	godotenv "github.com/joho/godotenv"
)

const (
	DEFAULT_OWN_TEAM = "HOGE"
)

func main() {
	inputExcelFilePath := flag.String("input-file", "", "Path to the Input Excel file")
	outputExcelFilePath := flag.String("output-file", "", "Path to the Output Excel file")
	flag.Parse()

	if *inputExcelFilePath == "" {
		log.Fatal("Input Excel file path is required")
	}

	ownTeam := DEFAULT_OWN_TEAM
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		ownTeam = os.Getenv("OWN_TEAM")
	}

	gameRepo := irepository.NewExcelGameRepository(*inputExcelFilePath, ownTeam)
	playerRepo := irepository.NewExcelPlayerRepository(*inputExcelFilePath)
	var rankingRepo drepository.RankingRepository
	if *outputExcelFilePath == "" {
		rankingRepo = irepository.NewStdoutRankingRepository()
	} else {
		rankingRepo = irepository.NewExcelRankingRepository(*outputExcelFilePath)
	}
	rg, err := usecase.NewRankingGenerator(playerRepo, gameRepo, rankingRepo)
	if err != nil {
		log.Fatal(err)
	}
	rg.PrintRankings()
}
