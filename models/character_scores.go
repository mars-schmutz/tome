package models

import "math"

type Ability int

const (
	Str Ability = iota
	Dex
	Con
	Int
	Wis
	Cha
)

type CharScores struct {
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

func (cs *CharScores) SetStrength(str int) {
	cs.Str = str
}

func (cs *CharScores) SetDexterity(dex int) {
	cs.Dex = dex
}

func (cs *CharScores) SetConstitution(con int) {
	cs.Con = con
}

func (cs *CharScores) SetIntelligence(intelli int) {
	cs.Int = intelli
}

func (cs *CharScores) SetWisdom(wis int) {
	cs.Wis = wis
}

func (cs *CharScores) SetCharisma(cha int) {
	cs.Cha = cha
}
