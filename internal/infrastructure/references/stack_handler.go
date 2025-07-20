package references

import (
	pulumiDynamodb "github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dynamodb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"pulumi-boardgame-assistant/internal/config"
	"pulumi-boardgame-assistant/internal/services"
)

type ReferencesStackHandler struct {
	referenceSeeder *services.ReferenceSeeder
	config          *config.Config
}

func NewReferencesStackHandler(referenceSeeder *services.ReferenceSeeder, cfg *config.Config) *ReferencesStackHandler {
	return &ReferencesStackHandler{
		referenceSeeder: referenceSeeder,
		config:          cfg,
	}
}

func (h *ReferencesStackHandler) CreateStack(ctx *pulumi.Context) error {
	table, err := h.createGameReferencesTable(ctx)
	if err != nil {
		return err
	}

	ctx.Export("referencesTableName", table.Name)
	ctx.Export("referencesTableArn", table.Arn)

	err = h.referenceSeeder.CreateReferenceResources(ctx, table)
	if err != nil {
		return err
	}

	return nil
}

func (h *ReferencesStackHandler) createGameReferencesTable(ctx *pulumi.Context) (*pulumiDynamodb.Table, error) {
	table, err := pulumiDynamodb.NewTable(ctx, "game-references", &pulumiDynamodb.TableArgs{
		Name: pulumi.String(h.config.DynamoDB.GameReferencesTable),
		Attributes: pulumiDynamodb.TableAttributeArray{
			&pulumiDynamodb.TableAttributeArgs{
				Name: pulumi.String("gameId"),
				Type: pulumi.String("S"),
			},
			&pulumiDynamodb.TableAttributeArgs{
				Name: pulumi.String("referenceId"),
				Type: pulumi.String("S"),
			},
		},
		HashKey:     pulumi.String("gameId"),
		RangeKey:    pulumi.String("referenceId"),
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
