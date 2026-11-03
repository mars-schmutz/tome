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

func (api *CharacterApi) SetName(name string) (models.Character, error) {
	c, err := api.svc.SetName(name)
	if err != nil {
		log.Printf("SetName(%q): %v", name, err)
	}
	return c, err
}

func (api *CharacterApi) GetCharacter() models.Character {
	return api.svc.GetCharacter()
}

func (api *CharacterApi) SetClass(class string) (models.Character, error) {
	c, err := api.svc.SetClass(class)
	if err != nil {
		log.Printf("SetClass(%q): %v", class, err)
	}
	return c, err
}

func (api *CharacterApi) SetRace(race string) (models.Character, error) {
	c, err := api.svc.SetRace(race)
	if err != nil {
		log.Printf("SetRace(%q): %v", race, err)
	}
	return c, err
}

func (api *CharacterApi) SetBackground(background string) (models.Character, error) {
	c, err := api.svc.SetBackground(background)
	if err != nil {
		log.Printf("SetBackground(%q): %v", background, err)
	}
	return c, err
}

func (api *CharacterApi) SetAlignment(alignment string) (models.Character, error) {
	c, err := api.svc.SetAlignment(alignment)
	if err != nil {
		log.Printf("SetAlignment(%q): %v", alignment, err)
	}
	return c, err
}

func (api *CharacterApi) SetLevel(level int) (models.Character, error) {
	c, err := api.svc.SetLevel(level)
	if err != nil {
		log.Printf("SetLevel(%d): %v", level, err)
	}
	return c, err
}

func (api *CharacterApi) SetInspiration(insp int) (models.Character, error) {
	c, err := api.svc.SetInspiration(insp)
	if err != nil {
		log.Printf("SetInspiration(%d): %v", insp, err)
	}
	return c, err
}

func (api *CharacterApi) SetScores(scores models.CharScores) (models.Character, error) {
	c, err := api.svc.SetScores(scores)
	if err != nil {
		log.Printf("SetScores(%+v): %v", scores, err)
	}
	return c, err
}

// SetSkillLevel specifically for setting the Skill.Level (0 none, 1 proficient, 2 expertise). Bonuses are calculated by models.Character
func (api *CharacterApi) SetSkillLevel(skill string, level int) (models.Character, error) {
	c, err := api.svc.SetSkillLevel(skill, level)
	if err != nil {
		log.Printf("SetSkillLevel(%q, %d): %v", skill, level, err)
	}
	return c, err
}

func (api *CharacterApi) SetSpellcastingAbility(ability models.Ability) (models.Character, error) {
	c, err := api.svc.SetSpellcastingAbility(ability)
	if err != nil {
		log.Printf("SetSpellcastingAbility(%q): %v", ability, err)
	}
	return c, err
}

func (api *CharacterApi) SetMaxHp(newHp int) (models.Character, error) {
	c, err := api.svc.SetMaxHp(newHp)
	if err != nil {
		log.Printf("SetMaxHp(%d): %v", newHp, err)
	}
	return c, err
}

func (api *CharacterApi) Heal(amount int) (models.Character, error) {
	c, err := api.svc.Heal(amount)
	if err != nil {
		log.Printf("Heal(%d): %v", amount, err)
	}
	return c, err
}

func (api *CharacterApi) TakeDamage(amount int) (models.Character, error) {
	c, err := api.svc.TakeDamage(amount)
	if err != nil {
		log.Printf("TakeDamage(%d): %v", amount, err)
	}
	return c, err
}

func (api *CharacterApi) SetHitDice(dice string) (models.Character, error) {
	c, err := api.svc.SetHitDice(dice)
	if err != nil {
		log.Printf("SetHitDice(%q): %v", dice, err)
	}
	return c, err
}

func (api *CharacterApi) SetArmorClass(ac int) (models.Character, error) {
	c, err := api.svc.SetArmorClass(ac)
	if err != nil {
		log.Printf("SetArmorClass(%d): %v", ac, err)
	}
	return c, err
}

func (api *CharacterApi) SetInitiative(initiative int) (models.Character, error) {
	c, err := api.svc.SetInitiative(initiative)
	if err != nil {
		log.Printf("SetInitiative(%d): %v", initiative, err)
	}
	return c, err
}

func (api *CharacterApi) SetSpeed(speed int) (models.Character, error) {
	c, err := api.svc.SetSpeed(speed)
	if err != nil {
		log.Printf("SetSpeed(%d): %v", speed, err)
	}
	return c, err
}

func (api *CharacterApi) SetDeathSaves(saves int) (models.Character, error) {
	c, err := api.svc.SetDeathSaves(saves)
	if err != nil {
		log.Printf("SetDeathSaves(%d): %v", saves, err)
	}
	return c, err
}

func (api *CharacterApi) SetDeathFails(fails int) (models.Character, error) {
	c, err := api.svc.SetDeathFails(fails)
	if err != nil {
		log.Printf("SetDeathFails(%d): %v", fails, err)
	}
	return c, err
}
