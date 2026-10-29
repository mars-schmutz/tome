// Package character holds the character service that handles race conditions and acts as middleman between api and models
package character

import (
	"sync"

	"tome/models"
	"tome/store"
)

type CharacterService struct {
	mu      sync.Mutex
	store   *store.Store
	current models.Character
}

func NewCharacterService(st *store.Store) *CharacterService {
	return &CharacterService{
		store:   st,
		current: models.NewCharacter(),
	}
}

func (s *CharacterService) GetCharacter() models.Character {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.current
}
