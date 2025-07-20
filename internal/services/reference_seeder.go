package services

import (
	"fmt"
	"log"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dynamodb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"pulumi-boardgame-assistant/internal/config"
	"pulumi-boardgame-assistant/internal/models"
)

type ReferenceSeeder struct {
	config *config.DynamoDB
}

func NewReferenceSeeder(cfg *config.DynamoDB) *ReferenceSeeder {
	log.Printf("Initializing reference seeder with table: %s", cfg.GameReferencesTable)

	return &ReferenceSeeder{
		config: cfg,
	}
}

func (rs *ReferenceSeeder) CreateReferenceResources(ctx *pulumi.Context, table *dynamodb.Table) error {
	references := models.GetReferences()

	for _, reference := range references {
		if err := reference.Validate(); err != nil {
			return fmt.Errorf("invalid reference %s: %w", reference.ReferenceId, err)
		}

		itemAttributes := fmt.Sprintf(`{
			"gameId": {"S": "%s"},
			"referenceId": {"S": "%s"},
			"type": {"S": "%s"},
			"title": {"S": "%s"},
			"section": {"S": "%s"},
			"pageReference": {"S": "%s"},
			"url": {"S": "%s"}
		}`, reference.GameId, reference.ReferenceId, reference.Type, reference.Title, reference.Section, reference.PageReference, reference.URL)

		_, err := dynamodb.NewTableItem(ctx, fmt.Sprintf("reference-%s-%s", reference.GameId, reference.ReferenceId), &dynamodb.TableItemArgs{
			TableName: table.Name,
			HashKey:   table.HashKey,
			RangeKey:  table.RangeKey,
			Item:      pulumi.String(itemAttributes),
		}, pulumi.DependsOn([]pulumi.Resource{table}))

		if err != nil {
			return fmt.Errorf("failed to create reference item %s: %w", reference.ReferenceId, err)
		}

		log.Printf("Created Pulumi resource for reference: %s-%s", reference.GameId, reference.ReferenceId)
	}

	log.Printf("Successfully created %d reference resources", len(references))
	return nil
}
