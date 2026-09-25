package models

// Skill Level 0 = none, 1 = proficient, 2 = expertise
type Skill struct {
	Bonus int
	Level int
}

type skillEntry struct {
	Name    string
	Skill   *Skill
	Ability Ability
}

type CharSkills struct {
	Acrobatics     Skill
	AnimalHandling Skill
	Arcana         Skill
	Athletics      Skill
	Deception      Skill
	History        Skill
	Insight        Skill
	Intimidation   Skill
	Investigation  Skill
	Medicine       Skill
	Nature         Skill
	Perception     Skill
	Performance    Skill
	Persuasion     Skill
	Religion       Skill
	SleightOfHand  Skill
	Stealth        Skill
	Survival       Skill
}

func (cs *CharSkills) entries() []skillEntry {
	return []skillEntry{
		{"Acrobatics", &cs.Acrobatics, Dex},
		{"AnimalHandling", &cs.AnimalHandling, Wis},
		{"Arcana", &cs.Arcana, Int},
		{"Athletics", &cs.Athletics, Str},
		{"Deception", &cs.Deception, Cha},
		{"History", &cs.History, Int},
		{"Insight", &cs.Insight, Wis},
		{"Intimidation", &cs.Intimidation, Cha},
		{"Investigation", &cs.Investigation, Int},
		{"Medicine", &cs.Medicine, Wis},
		{"Nature", &cs.Nature, Int},
		{"Perception", &cs.Perception, Wis},
		{"Performance", &cs.Performance, Cha},
		{"Persuasion", &cs.Persuasion, Cha},
		{"Religion", &cs.Religion, Int},
		{"SleightOfHand", &cs.SleightOfHand, Dex},
		{"Stealth", &cs.Stealth, Dex},
		{"Survival", &cs.Survival, Wis},
	}
}
