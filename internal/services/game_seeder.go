package services

import (
	"fmt"
	"log"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dynamodb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"pulumi-boardgame-assistant/internal/config"
	"pulumi-boardgame-assistant/internal/models"
)

type GameSeeder struct {
	config *config.DynamoDB
}

func NewGameSeeder(cfg *config.DynamoDB) *GameSeeder {
	log.Printf("Initializing game seeder with table: %s", cfg.SupportedGamesTable)

	return &GameSeeder{
		config: cfg,
	}
}

func (gs *GameSeeder) CreateGameResources(ctx *pulumi.Context, table *dynamodb.Table) error {
	games := models.GetSampleGames()

	for _, game := range games {
		if err := game.Validate(); err != nil {
			return fmt.Errorf("invalid game %s: %w", game.GameId, err)
		}

		itemAttributes := fmt.Sprintf(`{
			"gameId": {"S": "%s"},
			"name": {"S": "%s"},
			"publisher": {"S": "%s"},
			"year": {"N": "%d"}
		}`, game.GameId, game.Name, game.Publisher, game.Year)

		_, err := dynamodb.NewTableItem(ctx, fmt.Sprintf("game-%s", game.GameId), &dynamodb.TableItemArgs{
			TableName: table.Name,
			HashKey:   table.HashKey,
			Item:      pulumi.String(itemAttributes),
		}, pulumi.DependsOn([]pulumi.Resource{table}))

		if err != nil {
			return fmt.Errorf("failed to create game item %s: %w", game.GameId, err)
		}

		log.Printf("Created Pulumi resource for game: %s", game.GameId)
	}

	log.Printf("Successfully created %d game resources", len(games))
	return nil
}
