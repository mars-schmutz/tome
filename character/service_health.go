package character

import "tome/models"

func (s *CharacterService) SetMaxHp(newHp int) (models.Character, error) {
	return s.update("set max hp", func(c *models.Character) error {
		c.Health.SetMaxHp(newHp)
		return nil
	})
}

func (s *CharacterService) Heal(amount int) (models.Character, error) {
	return s.update("heal", func(c *models.Character) error {
		c.Health.Heal(amount)
		return nil
	})
}

func (s *CharacterService) TakeDamage(amount int) (models.Character, error) {
	return s.update("take damage", func(c *models.Character) error {
		c.Health.TakeDamage(amount)
		return nil
	})
}

func (s *CharacterService) SetHitDice(dice string) (models.Character, error) {
	return s.update("set hit dice", func(c *models.Character) error {
		c.Health.SetHitDice(dice)
		return nil
	})
}

func (s *CharacterService) SetArmorClass(ac int) (models.Character, error) {
	return s.update("set armor class", func(c *models.Character) error {
		c.Health.SetArmorClass(ac)
		return nil
	})
}

func (s *CharacterService) SetInitiative(initiative int) (models.Character, error) {
	return s.update("set initiative", func(c *models.Character) error {
		c.Health.SetInitiative(initiative)
		return nil
	})
}

func (s *CharacterService) SetSpeed(speed int) (models.Character, error) {
	return s.update("set speed", func(c *models.Character) error {
		c.Health.SetSpeed(speed)
		return nil
	})
}

func (s *CharacterService) SetDeathSaves(saves int) (models.Character, error) {
	return s.update("set death saves", func(c *models.Character) error {
		c.Health.SetDeathSaves(saves)
		return nil
	})
}

func (s *CharacterService) SetDeathFails(fails int) (models.Character, error) {
	return s.update("set death fails", func(c *models.Character) error {
		c.Health.SetDeathFails(fails)
		return nil
	})
}
