package main

import (
	"flag"
	"fmt"
	"it-league-stats/infrastructure/repository"
	"it-league-stats/usecase"
	"log"
	"os"

	godotenv "github.com/joho/godotenv"
	"github.com/kr/pretty"
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
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		ownTeam = os.Getenv("OWN_TEAM")
		log.Print("OWN_TEAM: ", ownTeam)
	}

	gameRepo := repository.NewExcelGameRepository(*excelFile, ownTeam)
	playerRepo := repository.NewExcelPlayerRepository(*excelFile)
	games, players, err := usecase.AllData(gameRepo, playerRepo)
	if err != nil {
		log.Fatal("Error loading Excel file")
	}
	for _, player := range players {
		fmt.Printf("Player: %s#%s\n", player.Name, player.ID)
	}
	fmt.Printf("Game: %# v", pretty.Formatter(games))
}
