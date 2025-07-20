package games

import (
	pulumiDynamodb "github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dynamodb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"pulumi-boardgame-assistant/internal/config"
	"pulumi-boardgame-assistant/internal/services"
)

type GamesStackHandler struct {
	gameSeeder *services.GameSeeder
	config     *config.Config
}

func NewGamesStackHandler(gameSeeder *services.GameSeeder, cfg *config.Config) *GamesStackHandler {
	return &GamesStackHandler{
		gameSeeder: gameSeeder,
		config:     cfg,
	}
}

func (h *GamesStackHandler) CreateStack(ctx *pulumi.Context) error {
	table, err := h.createSupportedGamesTable(ctx)
	if err != nil {
		return err
	}

	ctx.Export("gamesTableName", table.Name)
	ctx.Export("gamesTableArn", table.Arn)

	err = h.gameSeeder.CreateGameResources(ctx, table)
	if err != nil {
		return err
	}

	return nil
}

func (h *GamesStackHandler) createSupportedGamesTable(ctx *pulumi.Context) (*pulumiDynamodb.Table, error) {
	table, err := pulumiDynamodb.NewTable(ctx, "supported-games", &pulumiDynamodb.TableArgs{
		Name: pulumi.String(h.config.DynamoDB.SupportedGamesTable),
		Attributes: pulumiDynamodb.TableAttributeArray{
			&pulumiDynamodb.TableAttributeArgs{
				Name: pulumi.String("gameId"),
				Type: pulumi.String("S"),
			},
		},
		HashKey:     pulumi.String("gameId"),
		BillingMode: pulumi.String("PAY_PER_REQUEST"),
		PointInTimeRecovery: &pulumiDynamodb.TablePointInTimeRecoveryArgs{
			Enabled: pulumi.Bool(true),
		},
		ServerSideEncryption: &pulumiDynamodb.TableServerSideEncryptionArgs{
			Enabled: pulumi.Bool(true),
		},
	})

	if err != nil {
		return nil, err
	}

	return table, nil
}
