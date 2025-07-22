package models

import (
	"errors"
	"strings"
)

type Reference struct {
	GameId        string `dynamodbav:"gameId"`
	ReferenceId   string `dynamodbav:"referenceId"`
	Type          string `dynamodbav:"type"`
	Title         string `dynamodbav:"title"`
	Section       string `dynamodbav:"section"`
	PageReference string `dynamodbav:"pageReference"`
	URL           string `dynamodbav:"url"`
}

func (r *Reference) Validate() error {
	if strings.TrimSpace(r.GameId) == "" {
		return errors.New("gameId cannot be empty")
	}

	if strings.TrimSpace(r.ReferenceId) == "" {
		return errors.New("referenceId cannot be empty")
	}

	if strings.TrimSpace(r.Type) == "" {
		return errors.New("type cannot be empty")
	}

	if strings.TrimSpace(r.Title) == "" {
		return errors.New("title cannot be empty")
	}

	if strings.TrimSpace(r.Section) == "" {
		return errors.New("section cannot be empty")
	}

	return nil
}

func GetReferences() []Reference {
	return []Reference{
		{
			GameId:        "nemesis",
			ReferenceId:   "ATK1",
			Type:          "deck",
			Title:         "Nemesis Core Game",
			Section:       "Intruder Attack Deck",
			PageReference: "",
			URL:           "",
		},
		{
			GameId:        "nemesis",
			ReferenceId:   "CHARACTION",
			Type:          "deck",
			Title:         "Nemesis Core Game",
			Section:       "Character Action Decks",
			PageReference: "",
			URL:           "",
		},
		{
			GameId:        "nemesis",
			ReferenceId:   "EVT1",
			Type:          "deck",
			Title:         "Nemesis Core Game",
			Section:       "Event Deck",
			PageReference: "",
			URL:           "",
		},
		{
			GameId:        "nemesis",
			ReferenceId:   "ITM1-BLUE",
			Type:          "deck",
			Title:         "Nemesis Core Game",
			Section:       "Craft (Blue) Item Deck",
			PageReference: "",
			URL:           "",
		},
		{
			GameId:        "nemesis",
			ReferenceId:   "ITM1-GREEN",
			Type:          "deck",
			Title:         "Nemesis Core Game",
			Section:       "Medical (Green) Item Deck",
			PageReference: "",
			URL:           "",
		},
		{
			GameId:        "nemesis",
			ReferenceId:   "ITM1-RED",
			Type:          "deck",
			Title:         "Nemesis Core Game",
			Section:       "Military (Red) Item Deck",
			PageReference: "",
			URL:           "",
		},
		{
			GameId:        "nemesis",
			ReferenceId:   "ITM1-YELLOW",
			Type:          "deck",
			Title:         "Nemesis Core Game",
			Section:       "Technical (Yellow) Item Deck",
			PageReference: "",
			URL:           "",
		},
		{
			GameId:        "nemesis",
			ReferenceId:   "MECH1",
			Type:          "deck",
			Title:         "Nemesis Core Game",
			Section:       "Mechanic Action Deck",
			PageReference: "",
			URL:           "",
		},
		{
			GameId:        "nemesis",
			ReferenceId:   "MECH2",
			Type:          "deck",
			Title:         "Nemesis Core Game",
			Section:       "Mechanic Quest Items",
			PageReference: "",
			URL:           "",
		},
		{
			GameId:        "nemesis",
			ReferenceId:   "R1-DOORS",
			Type:          "rulebook",
			Title:         "Nemesis Core Rulebook",
			Section:       "Door Token",
			PageReference: "p.17",
			URL:           "",
		},
		{
			GameId:        "nemesis",
			ReferenceId:   "R1-FIRE",
			Type:          "rulebook",
			Title:         "Nemesis Core Rulebook",
			Section:       "Fire Markers",
			PageReference: "p.17",
			URL:           "",
		},
		{
			GameId:        "nemesis",
			ReferenceId:   "R1-MOVEMENT",
			Type:          "rulebook",
			Title:         "Nemesis Core Rulebook",
			Section:       "Movement and Exploration",
			PageReference: "p.15",
			URL:           "",
		},
		{
			GameId:        "nemesis",
			ReferenceId:   "R1-NOISE",
			Type:          "rulebook",
			Title:         "Nemesis Core Rulebook",
			Section:       "Noise Roll",
			PageReference: "p.15",
			URL:           "",
		},
		{
			GameId:        "nemesis",
			ReferenceId:   "R1-ROOMS",
			Type:          "rulebook",
			Title:         "Nemesis Core Rulebook",
			Section:       "Room Sheets",
			PageReference: "p.25",
			URL:           "",
		},
		{
			GameId:        "nemesis",
			ReferenceId:   "R1-SLIME",
			Type:          "rulebook",
			Title:         "Nemesis Core Rulebook",
			Section:       "Slime Markers",
			PageReference: "p.17",
			URL:           "",
		},
	}
}
