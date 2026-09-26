package main

import (
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
