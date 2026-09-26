// Package character holds the character service that handles race conditions and acts as middleman between api and models
package character

import (
	"sync"

	"tome/models"
)

type CharacterService struct {
	mu      sync.Mutex
	current models.Character
}

func NewCharacterService() *CharacterService {
	return &CharacterService{
		current: models.NewCharacter(),
	}
}

func (s *CharacterService) SetName(name string) models.Character {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.current.Base.SetName(name)
	return s.current
}

func (s *CharacterService) GetCharacter() models.Character {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.current
}
