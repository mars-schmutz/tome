package main

import (
	"log"

	"tome/character"
	"tome/models"
)

type CharacterApi struct {
	svc *character.CharacterService
}

func NewCharacterApi(newSvc *character.CharacterService) *CharacterApi {
	return &CharacterApi{
		svc: newSvc,
	}
}

func (api *CharacterApi) SetName(name string) models.Character {
	return api.svc.SetName(name)
}

func (api *CharacterApi) GetCharacter() models.Character {
	return api.svc.GetCharacter()
}

func (api *CharacterApi) SetClass(class string) models.Character {
	return api.svc.SetClass(class)
}

func (api *CharacterApi) SetRace(race string) models.Character {
	return api.svc.SetRace(race)
}

func (api *CharacterApi) SetBackground(background string) models.Character {
	return api.svc.SetBackground(background)
}

func (api *CharacterApi) SetAlignment(alignment string) models.Character {
	return api.svc.SetAlignment(alignment)
}

func (api *CharacterApi) SetLevel(level int) models.Character {
	return api.svc.SetLevel(level)
}

func (api *CharacterApi) SetInspiration(insp int) models.Character {
	return api.svc.SetInspiration(insp)
}

func (api *CharacterApi) SetScores(scores models.CharScores) models.Character {
	return api.svc.SetScores(scores)
}

// SetSkillLevel specifically for setting the Skill.Level (0 none, 1 proficient, 2 expertise). Bonuses are calculated by models.Character
func (api *CharacterApi) SetSkillLevel(skill string, level int) (models.Character, error) {
	c, err := api.svc.SetSkillLevel(skill, level)
	if err != nil {
		log.Printf("SetSkillLevel(%q, %d): %v", skill, level, err)
	}
	return c, err
}

func (api *CharacterApi) SetMaxHp(newHp int) models.Character {
	return api.svc.SetMaxHp(newHp)
}

func (api *CharacterApi) Heal(amount int) models.Character {
	return api.svc.Heal(amount)
}

func (api *CharacterApi) TakeDamage(amount int) models.Character {
	return api.svc.TakeDamage(amount)
}

func (api *CharacterApi) SetHitDice(dice string) models.Character {
	return api.svc.SetHitDice(dice)
}

func (api *CharacterApi) SetArmorClass(ac int) models.Character {
	return api.svc.SetArmorClass(ac)
}

func (api *CharacterApi) SetInitiative(initiative int) models.Character {
	return api.svc.SetInitiative(initiative)
}

func (api *CharacterApi) SetSpeed(speed int) models.Character {
	return api.svc.SetSpeed(speed)
}

func (api *CharacterApi) SetDeathSaves(saves int) models.Character {
	return api.svc.SetDeathSaves(saves)
}

func (api *CharacterApi) SetDeathFails(fails int) models.Character {
	return api.svc.SetDeathFails(fails)
}
