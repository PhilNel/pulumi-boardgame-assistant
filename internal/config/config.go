package config

import (
	"fmt"

	"github.com/jessevdk/go-flags"
)

type Config struct {
	DynamoDB *DynamoDB
}

type DynamoDB struct {
	SupportedGamesTable string `long:"supported_games_table" env:"SUPPORTED_GAMES_TABLE" description:"DynamoDB table for supported games" default:"boardgame-assistant-supported-games-dev"`
	GameReferencesTable string `long:"game_references_table" env:"GAME_REFERENCES_TABLE" description:"DynamoDB table for game references" default:"boardgame-assistant-game-references-dev"`
}

func Load() (*Config, error) {
	opts := &Config{}
	_, err := flags.Parse(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return opts, nil
}
