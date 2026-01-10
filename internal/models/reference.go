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

// GetReferences aggregates all references from all games
func GetReferences() []Reference {
	var refs []Reference
	refs = append(refs, getNemesisCharacterReferences()...)
	refs = append(refs, getNemesisDeckReferences()...)
	refs = append(refs, getNemesisRulebookReferences()...)
	return refs
}
