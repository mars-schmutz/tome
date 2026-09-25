package models

import "testing"

func TestSetSkillLevel(t *testing.T) {
	tests := []struct {
		name      string
		skill     string
		level     int
		wantErr   bool
		wantLevel int
		wantBonus int
	}{
		{"set expertise", "Stealth", 2, false, 2, 6},
		{"clear to none", "Stealth", 0, false, 0, 2},
		{"unknown skill", "Stelth", 1, true, 0, 0},
		{"empty skill name", "", 1, true, 0, 0},
		{"names are case sensitive", "stealth", 1, true, 0, 0},
		{"level too high", "Stealth", 3, true, 0, 0},
		{"level negative", "Stealth", -1, true, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Start every case from the same known state: proficient Stealth, bonus 4.
			c := Character{Scores: CharScores{Dex: 14}, Proficiency: 2}
			c.Skills.Stealth.Level = 1
			c.RecalculateAllSkills()
			before := c.Skills

			err := c.SetSkillLevel(tt.skill, tt.level)

			if (err != nil) != tt.wantErr {
				t.Fatalf("SetSkillLevel(%q, %d) error = %v, wantErr %v", tt.skill, tt.level, err, tt.wantErr)
			}

			if tt.wantErr {
				if c.Skills != before {
					t.Errorf("skills changed on a failed call:\n got  %+v\n want %+v", c.Skills, before)
				}
				return
			}

			if got := c.Skills.Stealth.Level; got != tt.wantLevel {
				t.Errorf("Stealth Level = %d, want %d", got, tt.wantLevel)
			}
			if got := c.Skills.Stealth.Bonus; got != tt.wantBonus {
				t.Errorf("Stealth Bonus = %d, want %d", got, tt.wantBonus)
			}
		})
	}
}
