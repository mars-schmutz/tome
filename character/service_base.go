package character

import (
	"tome/models"
)

func (s *CharacterService) SetName(name string) (models.Character, error) {
	return s.update("set name", func(c *models.Character) error {
		c.Base.SetName(name)
		return nil
	})
}

func (s *CharacterService) SetClass(class string) (models.Character, error) {
	return s.update("set class", func(c *models.Character) error {
		c.Base.SetClass(class)
		return nil
	})
}

func (s *CharacterService) SetRace(race string) (models.Character, error) {
	return s.update("set race", func(c *models.Character) error {
		c.Base.SetRace(race)
		return nil
	})
}

func (s *CharacterService) SetBackground(background string) (models.Character, error) {
	return s.update("set background", func(c *models.Character) error {
		c.Base.SetBackground(background)
		return nil
	})
}

func (s *CharacterService) SetAlignment(alignment string) (models.Character, error) {
	return s.update("set alignment", func(c *models.Character) error {
		c.Base.SetAlignment(alignment)
		return nil
	})
}

func (s *CharacterService) SetInspiration(insp int) (models.Character, error) {
	return s.update("set inspiration", func(c *models.Character) error {
		c.SetInspiration(insp)
		return nil
	})
}

func (s *CharacterService) SetLevel(level int) (models.Character, error) {
	return s.update("set level", func(c *models.Character) error {
		c.SetLevel(level)
		return nil
	})
}

func (s *CharacterService) SetScores(scores models.CharScores) (models.Character, error) {
	return s.update("set scores", func(c *models.Character) error {
		c.SetScores(scores)
		return nil
	})
}

func (s *CharacterService) SetSkillLevel(skill string, level int) (models.Character, error) {
	return s.update("set skill level", func(c *models.Character) error {
		return c.SetSkillLevel(skill, level)
	})
}

func (s *CharacterService) SetSpellcastingAbility(ability models.Ability) (models.Character, error) {
	return s.update("set spellcasting ability", func(c *models.Character) error {
		return c.SetSpellcastingAbility(ability)
	})
}
