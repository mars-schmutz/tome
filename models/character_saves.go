package models

type Save struct {
	Proficient bool
	Bonus      int
}

type CharSaves struct {
	Str Save
	Dex Save
	Con Save
	Int Save
	Wis Save
	Cha Save
}

type saveEntry struct {
	Save    *Save
	Ability Ability
}

func (cs *CharSaves) entries() []saveEntry {
	return []saveEntry{
		{&cs.Str, Str},
		{&cs.Dex, Dex},
		{&cs.Con, Con},
		{&cs.Int, Int},
		{&cs.Wis, Wis},
		{&cs.Cha, Cha},
	}
}
