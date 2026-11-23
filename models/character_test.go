package models

import (
	"fmt"
	"testing"
)

func TestSetSkillTier(t *testing.T) {
	tests := []struct {
		name      string
		skill     string
		tier      int
		wantErr   bool
		wantTier  int
		wantBonus int
	}{
		{"set expertise", "Stealth", 2, false, 2, 6},
		{"clear to none", "Stealth", 0, false, 0, 2},
		{"unknown skill", "Stelth", 1, true, 0, 0},
		{"empty skill name", "", 1, true, 0, 0},
		{"names are case sensitive", "stealth", 1, true, 0, 0},
		{"tier too high", "Stealth", 3, true, 0, 0},
		{"tier negative", "Stealth", -1, true, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Start every case from the same known state: proficient Stealth, bonus 4.
			c := Character{Scores: CharScores{Dex: 14}, Proficiency: 2}
			c.Skills.Stealth.Tier = 1
			c.recalculateAllSkills()
			before := c.Skills

			err := c.SetSkillTier(tt.skill, tt.tier)

			if (err != nil) != tt.wantErr {
				t.Fatalf("SetSkillTier(%q, %d) error = %v, wantErr %v", tt.skill, tt.tier, err, tt.wantErr)
			}

			if tt.wantErr {
				if c.Skills != before {
					t.Errorf("skills changed on a failed call:\n got  %+v\n want %+v", c.Skills, before)
				}
				return
			}

			if got := c.Skills.Stealth.Tier; got != tt.wantTier {
				t.Errorf("Stealth Tier = %d, want %d", got, tt.wantTier)
			}
			if got := c.Skills.Stealth.Bonus; got != tt.wantBonus {
				t.Errorf("Stealth Bonus = %d, want %d", got, tt.wantBonus)
			}
		})
	}
}

func TestSetLevel(t *testing.T) {
	tests := []struct {
		level     int
		wantProf  int
		wantBonus int // proficient Stealth with Dex 14 (+2), so 2 + PB
	}{
		{1, 2, 4},
		{4, 2, 4},
		{5, 3, 5},
		{8, 3, 5},
		{9, 4, 6},
		{12, 4, 6},
		{13, 5, 7},
		{16, 5, 7},
		{17, 6, 8},
		{20, 6, 8},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("level %d", tt.level), func(t *testing.T) {
			c := Character{Scores: CharScores{Dex: 14}}
			c.Skills.Stealth.Tier = 1

			c.SetLevel(tt.level)

			if c.Base.Level != tt.level {
				t.Errorf("Base.Level = %d, want %d", c.Base.Level, tt.level)
			}
			if c.Proficiency != tt.wantProf {
				t.Errorf("Proficiency = %d, want %d", c.Proficiency, tt.wantProf)
			}
			// Make sure recalculate ran when changing level (SetLevel -> recalculate())
			if got := c.Skills.Stealth.Bonus; got != tt.wantBonus {
				t.Errorf("Stealth Bonus = %d, want %d", got, tt.wantBonus)
			}
		})
	}
}

func TestSetScores(t *testing.T) {
	tests := []struct {
		name   string
		scores CharScores
		skill  string
		tier   int
		want   int
	}{
		{"raising Dex updates a proficient Dex skill", CharScores{Str: 10, Dex: 16, Con: 10, Int: 10, Wis: 10, Cha: 10}, "Acrobatics", 1, 5},
		{"lowering Wis updates an unproficient Wis skill", CharScores{Str: 10, Dex: 10, Con: 10, Int: 10, Wis: 7, Cha: 10}, "Perception", 0, -2},
		{"expertise with a high Int", CharScores{Str: 10, Dex: 10, Con: 10, Int: 20, Wis: 10, Cha: 10}, "Arcana", 2, 9},
		{"missing fields become 0", CharScores{Dex: 16}, "Athletics", 0, -5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Start from a flat 10 in everything, proficiency 2, so only new scores matter
			c := Character{Proficiency: 2}
			c.SetScores(CharScores{Str: 10, Dex: 10, Con: 10, Int: 10, Wis: 10, Cha: 10})
			if err := c.SetSkillTier(tt.skill, tt.tier); err != nil {
				t.Fatalf("setup: SetSkillTier(%q, %d): %v", tt.skill, tt.tier, err)
			}

			c.SetScores(tt.scores)

			if c.Scores != tt.scores {
				t.Errorf("Scores = %+v, want %+v", c.Scores, tt.scores)
			}
			if got := findSkill(t, &c.Skills, tt.skill).Bonus; got != tt.want {
				t.Errorf("%s Bonus = %d, want %d", tt.skill, got, tt.want)
			}
		})
	}
}
