package models

import (
	"errors"
	"strings"
)

type Game struct {
	GameId    string `dynamodbav:"gameId"`
	Name      string `dynamodbav:"name"`
	Publisher string `dynamodbav:"publisher"`
	Year      int    `dynamodbav:"year"`
}

func (g *Game) Validate() error {
	if strings.TrimSpace(g.GameId) == "" {
		return errors.New("gameId cannot be empty")
	}

	if strings.TrimSpace(g.Name) == "" {
		return errors.New("name cannot be empty")
	}

	if strings.TrimSpace(g.Publisher) == "" {
		return errors.New("publisher cannot be empty")
	}

	if g.Year < 1900 || g.Year > 2050 {
		return errors.New("year must be between 1900 and 2050")
	}

	return nil
}

func GetGames() []Game {
	return []Game{
		{
			GameId:    "nemesis",
			Name:      "Nemesis",
			Publisher: "Awaken Realms",
			Year:      2018,
		},
	}
}
