package character

import (
	"tome/models"
)

func (s *CharacterService) SetName(name string) models.Character {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.current.Base.SetName(name)
	return s.current
}

func (s *CharacterService) SetClass(class string) models.Character {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.current.Base.SetClass(class)
	return s.current
}

func (s *CharacterService) SetRace(race string) models.Character {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.current.Base.SetRace(race)
	return s.current
}

func (s *CharacterService) SetBackground(background string) models.Character {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.current.Base.SetBackground(background)
	return s.current
}

func (s *CharacterService) SetAlignment(alignment string) models.Character {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.current.Base.SetAlignment(alignment)
	return s.current
}

func (s *CharacterService) SetInspiration(insp int) models.Character {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.current.SetInspiration(insp)
	return s.current
}

func (s *CharacterService) SetLevel(level int) models.Character {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.current.SetLevel(level)
	return s.current
}

func (s *CharacterService) SetScores(scores models.CharScores) models.Character {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.current.SetScores(scores)
	return s.current
}
