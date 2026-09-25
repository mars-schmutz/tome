package models

import "testing"

func TestHeal(t *testing.T) {
	h := CharHealth{MaxHp: 10, CurrHp: 5}
	h.Heal(3)
	if h.CurrHp != 8 {
		t.Errorf("CurrHp = %d, want 8", h.CurrHp)
	}
}

// helper to make sure skills are all named properly
func findSkill(t *testing.T, cs *CharSkills, name string) *Skill {
	t.Helper()
	for _, e := range cs.entries() {
		if e.Name == name {
			return e.Skill
		}
	}
	t.Fatalf("no skill named %q in entries()", name)
	return nil
}

func TestRecalculateAllSkills(t *testing.T) {
	tests := []struct {
		name        string
		scores      CharScores
		proficiency int
		skill       string
		level       int
		want        int
	}{
		{"proficient", CharScores{Dex: 14}, 2, "Acrobatics", 1, 4},
		{"not proficient, odd score below 10", CharScores{Str: 9}, 2, "Athletics", 0, -1},
		{"expertise", CharScores{Dex: 16}, 3, "Stealth", 2, 9},
		{"proficient with negative mod", CharScores{Cha: 8}, 2, "Persuasion", 1, 1},
		{"Insight uses Wis, not Int", CharScores{Wis: 18, Int: 10}, 2, "Insight", 0, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Character{Scores: tt.scores, Proficiency: tt.proficiency}
			findSkill(t, &c.Skills, tt.skill).Level = tt.level

			c.RecalculateAllSkills()

			if got := findSkill(t, &c.Skills, tt.skill).Bonus; got != tt.want {
				t.Errorf("%s Bonus = %d, want %d", tt.skill, got, tt.want)
			}
		})
	}
}
