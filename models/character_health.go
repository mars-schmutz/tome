package models

type CharHealth struct {
	MaxHp      int
	CurrHp     int
	TempHp     int
	HitDice    string
	ArmorClass int
	Initiative int
	Speed      int
	DeathSaves int
	DeathFails int
}

func (ch *CharHealth) SetMaxHp(newHp int) {
	ch.MaxHp = newHp
}

func (ch *CharHealth) Heal(amount int) {
	ch.CurrHp += amount
	if ch.CurrHp > ch.MaxHp {
		ch.CurrHp = ch.MaxHp
	}
}

func (ch *CharHealth) TakeDamage(amount int) {
	if ch.TempHp > 0 {
		ch.TempHp -= amount
		if ch.TempHp < 0 {
			ch.CurrHp += ch.TempHp
			ch.TempHp = 0
		}
	} else {
		ch.CurrHp -= amount
	}

	if ch.CurrHp < 0 {
		ch.CurrHp = 0
	}
}

func (ch *CharHealth) SetHitDice(dice string) {
	ch.HitDice = dice
}

func (ch *CharHealth) SetArmorClass(ac int) {
	ch.ArmorClass = ac
}

func (ch *CharHealth) SetInitiative(initiative int) {
	ch.Initiative = initiative
}

func (ch *CharHealth) SetSpeed(speed int) {
	ch.Speed = speed
}

func (ch *CharHealth) SetDeathSaves(saves int) {
	ch.DeathSaves = saves
}

func (ch *CharHealth) SetDeathFails(fails int) {
	ch.DeathFails = fails
}
