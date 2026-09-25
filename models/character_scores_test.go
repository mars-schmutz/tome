package models

import "testing"

func TestAbilityModifier(t *testing.T) {
	tests := []struct {
		name    string
		scores  CharScores
		ability Ability
		want    int
	}{
		// odd scores below 10 round down not toward zero
		{"score 1", CharScores{Str: 1}, Str, -5},
		{"score 8", CharScores{Str: 8}, Str, -1},
		{"score 9", CharScores{Str: 9}, Str, -1},
		{"score 10", CharScores{Str: 10}, Str, 0},
		{"score 11", CharScores{Str: 11}, Str, 0},
		{"score 12", CharScores{Str: 12}, Str, 1},
		{"score 20", CharScores{Str: 20}, Str, 5},

		// each ability reads its own field not a neighbor's
		{"reads Dex", CharScores{Dex: 14}, Dex, 2},
		{"reads Con", CharScores{Con: 16}, Con, 3},
		{"reads Int", CharScores{Int: 18}, Int, 4},
		{"reads Wis", CharScores{Wis: 6}, Wis, -2},
		{"reads Cha", CharScores{Cha: 20}, Cha, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.scores.AbilityModifier(tt.ability); got != tt.want {
				t.Errorf("AbilityModifier(%v) = %d, want %d", tt.ability, got, tt.want)
			}
		})
	}
}
