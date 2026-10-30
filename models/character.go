// Package models provides all the structures for the data
package models

import (
	"encoding/json"
	"fmt"
)

type CharacterSummary struct {
	ID   int64
	Name string
}

type CharBase struct {
	Name       string
	Class      string
	Race       string
	Background string
	Alignment  string
	Level      int
}

type Character struct {
	ID             int64
	Base           CharBase
	Scores         CharScores
	ScoreModifiers CharModifiers
	Health         CharHealth
	Skills         CharSkills
	Spellcasting   CharSpellcasting
	Inspiration    int
	Proficiency    int
}

func NewCharacter() Character {
	char := Character{
		Base:         CharBase{},
		Scores:       CharScores{},
		Health:       CharHealth{},
		Skills:       CharSkills{},
		Spellcasting: CharSpellcasting{},
		Inspiration:  0,
		Proficiency:  0,
	}
	char.SetLevel(1)
	return char
}

func DecodeCharacter(data []byte) (Character, error) {
	var c Character
	err := json.Unmarshal(data, &c)
	if err != nil {
		return Character{}, fmt.Errorf("decode character: %w", err)
	}
	c.recalculate()
	return c, nil
}

func (cb *CharBase) SetName(name string) {
	cb.Name = name
}

func (cb *CharBase) SetClass(class string) {
	cb.Class = class
}

func (cb *CharBase) SetRace(race string) {
	cb.Race = race
}

func (cb *CharBase) SetBackground(background string) {
	cb.Background = background
}

func (cb *CharBase) SetAlignment(alignment string) {
	cb.Alignment = alignment
}

func (c *Character) SetInspiration(insp int) {
	c.Inspiration = insp
}

func (c *Character) SetLevel(level int) {
	c.Base.Level = level
	c.recalculate()
}

func (c *Character) SetScores(scores CharScores) {
	c.Scores = scores
	c.recalculate()
}

func (c *Character) recalculateProficiency() {
	if c.Base.Level <= 4 {
		c.Proficiency = 2
	} else if c.Base.Level <= 8 {
		c.Proficiency = 3
	} else if c.Base.Level <= 12 {
		c.Proficiency = 4
	} else if c.Base.Level <= 16 {
		c.Proficiency = 5
	} else {
		c.Proficiency = 6
	}
}

func (c *Character) SetSpellcastingAbility(ability Ability) error {
	// Empty string is for non-spellcasters
	switch ability {
	case Str, Dex, Con, Int, Wis, Cha, "":
		c.Spellcasting.Ability = ability
		c.recalculate()
		return nil
	default:
		return fmt.Errorf("invalid spellcasting ability: %q", ability)
	}
}

func (c *Character) recalculateSpellcasting() {
	if c.Spellcasting.Ability == "" {
		c.Spellcasting.AttackBonus = 0
		c.Spellcasting.SaveDC = 0
		return
	}
	abilityMod := c.Scores.AbilityModifier(c.Spellcasting.Ability)
	dc := 8 + c.Proficiency + abilityMod
	attackBonus := c.Proficiency + abilityMod

	c.Spellcasting.SaveDC = dc
	c.Spellcasting.AttackBonus = attackBonus
}

func (c *Character) recalculate() {
	c.recalculateProficiency()
	c.recalculateModifiers()
	c.recalculateAllSkills()
	c.recalculateSpellcasting()
}

func (c *Character) recalculateAllSkills() {
	for _, e := range c.Skills.entries() {
		skillBonus := e.Skill.Level * c.Proficiency
		modifier := c.Scores.AbilityModifier(e.Ability)
		e.Skill.Bonus = skillBonus + modifier
	}
}

func (c *Character) recalculateModifiers() {
	c.ScoreModifiers.Str = c.Scores.AbilityModifier(Str)
	c.ScoreModifiers.Dex = c.Scores.AbilityModifier(Dex)
	c.ScoreModifiers.Con = c.Scores.AbilityModifier(Con)
	c.ScoreModifiers.Int = c.Scores.AbilityModifier(Int)
	c.ScoreModifiers.Wis = c.Scores.AbilityModifier(Wis)
	c.ScoreModifiers.Cha = c.Scores.AbilityModifier(Cha)
}

func (c *Character) SetSkillLevel(skill string, level int) error {
	if level > 2 || level < 0 {
		return fmt.Errorf("invalid level value %d", level)
	}

	for _, e := range c.Skills.entries() {
		if e.Name == skill {
			e.Skill.Level = level
			c.recalculate()
			return nil
		}
	}

	return fmt.Errorf("unknown skill %q", skill)
}
