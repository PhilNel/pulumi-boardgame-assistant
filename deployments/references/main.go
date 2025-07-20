package main

import (
	"log"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"pulumi-boardgame-assistant/internal/config"
	"pulumi-boardgame-assistant/internal/infrastructure/references"
	"pulumi-boardgame-assistant/internal/services"
)

var referencesStackHandler *references.ReferencesStackHandler

func init() {
	log.Printf("Starting References Stack initialization")

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	referenceSeeder := services.NewReferenceSeeder(cfg.DynamoDB)

	referencesStackHandler = references.NewReferencesStackHandler(referenceSeeder, cfg)

	log.Printf("References Stack initialized successfully")
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("PANIC: Stack handler panicked: %v", r)
			}
		}()

		return referencesStackHandler.CreateStack(ctx)
	})
}
