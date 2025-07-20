package main

import (
	"log"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"pulumi-boardgame-assistant/internal/config"
	"pulumi-boardgame-assistant/internal/infrastructure/games"
	"pulumi-boardgame-assistant/internal/services"
)

var gamesStackHandler *games.GamesStackHandler

func init() {
	log.Printf("Starting Games Stack initialization")

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	gameSeeder := services.NewGameSeeder(cfg.DynamoDB)

	gamesStackHandler = games.NewGamesStackHandler(gameSeeder, cfg)

	log.Printf("Games Stack initialized successfully")
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("PANIC: Stack handler panicked: %v", r)
			}
		}()

		return gamesStackHandler.CreateStack(ctx)
	})
}
