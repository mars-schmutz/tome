// Package models provides all the structures for the data
package models

import "fmt"

type CharBase struct {
	Name       string
	Class      string
	Race       string
	Background string
	Alignment  string
	Level      int
}

type Character struct {
	ID          int64
	Base        CharBase
	Scores      CharScores
	Health      CharHealth
	Skills      CharSkills
	Inspiration int
	Proficiency int
}

func NewCharacter() Character {
	return Character{
		Base:        CharBase{},
		Scores:      CharScores{},
		Health:      CharHealth{},
		Skills:      CharSkills{},
		Inspiration: 0,
		Proficiency: 0,
	}
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
	c.recalculateProficiency()
}

func (c *Character) SetScores(scores CharScores) {
	c.Scores = scores
	c.recalculateAllSkills()
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
	c.recalculateAllSkills()
}

func (c *Character) recalculateAllSkills() {
	for _, e := range c.Skills.entries() {
		skillBonus := e.Skill.Level * c.Proficiency
		modifier := c.Scores.AbilityModifier(e.Ability)
		e.Skill.Bonus = skillBonus + modifier
	}
}

func (c *Character) SetSkillLevel(skill string, level int) error {
	if level > 2 || level < 0 {
		return fmt.Errorf("invalid level value %d", level)
	}

	for _, e := range c.Skills.entries() {
		if e.Name == skill {
			e.Skill.Level = level
			c.recalculateAllSkills()
			return nil
		}
	}

	return fmt.Errorf("unknown skill %q", skill)
}
