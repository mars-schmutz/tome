package character

import "tome/models"

func (s *CharacterService) SetMaxHp(newHp int) models.Character {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.current.Health.SetMaxHp(newHp)
	return s.current
}

func (s *CharacterService) Heal(amount int) models.Character {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.current.Health.Heal(amount)
	return s.current
}

func (s *CharacterService) TakeDamage(amount int) models.Character {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.current.Health.TakeDamage(amount)
	return s.current
}

func (s *CharacterService) SetHitDice(dice string) models.Character {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.current.Health.SetHitDice(dice)
	return s.current
}

func (s *CharacterService) SetArmorClass(ac int) models.Character {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.current.Health.SetArmorClass(ac)
	return s.current
}

func (s *CharacterService) SetInitiative(initiative int) models.Character {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.current.Health.SetInitiative(initiative)
	return s.current
}

func (s *CharacterService) SetSpeed(speed int) models.Character {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.current.Health.SetSpeed(speed)
	return s.current
}

func (s *CharacterService) SetDeathSaves(saves int) models.Character {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.current.Health.SetDeathSaves(saves)
	return s.current
}

func (s *CharacterService) SetDeathFails(fails int) models.Character {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.current.Health.SetDeathFails(fails)
	return s.current
}
