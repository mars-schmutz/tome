package models

import "math"

type Ability string

const (
	Str Ability = "Str"
	Dex Ability = "Dex"
	Con Ability = "Con"
	Int Ability = "Int"
	Wis Ability = "Wis"
	Cha Ability = "Cha"
)

type CharScores struct {
	Str int
	Dex int
	Con int
	Int int
	Wis int
	Cha int
}

type CharModifiers struct {
	Str int
	Dex int
	Con int
	Int int
	Wis int
	Cha int
}

func (cs *CharScores) AbilityModifier(ability Ability) int {
	var score int

	switch ability {
	case Str:
		score = cs.Str
	case Dex:
		score = cs.Dex
	case Con:
		score = cs.Con
	case Int:
		score = cs.Int
	case Wis:
		score = cs.Wis
	case Cha:
		score = cs.Cha
	default:
		return 0
	}

	return int(math.Floor(float64(score-10) / 2))
}
