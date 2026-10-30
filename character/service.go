// Package character holds the character service that handles race conditions and acts as middleman between api and models
package character

import (
	"errors"
	"fmt"
	"sync"

	"tome/models"
	"tome/store"
)

var ErrNoCharacter = errors.New("no character selected")

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

func (s *CharacterService) ListCharacters() ([]models.CharacterSummary, error) {
	list, err := s.store.ListCharacters()
	if err != nil {
		return nil, fmt.Errorf("list characters: %w", err)
	}

	return list, nil
}

func (s *CharacterService) CreateCharacter() (models.Character, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var c models.Character
	id, err := s.store.CreateCharacter(c)
	if err != nil {
		return models.Character{}, fmt.Errorf("create character: %w", err)
	}
	c.ID = id
	s.current = c
	return s.current, nil
}

func (s *CharacterService) SelectCharacter(id int64) (models.Character, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	c, err := s.store.LoadCharacter(id)
	if err != nil {
		return models.Character{}, fmt.Errorf("select character: %w", err)
	}
	s.current = c
	return s.current, nil
}

func (s *CharacterService) update(action string, mutate func(c *models.Character) error) (models.Character, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.current.ID == 0 {
		return models.Character{}, fmt.Errorf("%s: %w", action, ErrNoCharacter)
	}
	c := s.current
	if err := mutate(&c); err != nil {
		return models.Character{}, fmt.Errorf("%s: %w", action, ErrNoCharacter)
	}
	if err := s.store.SaveCharacter(c); err != nil {
		return models.Character{}, fmt.Errorf("%s: %w", action, ErrNoCharacter)
	}
	s.current = c
	return s.current, nil
}
