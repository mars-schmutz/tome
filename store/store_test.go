package store

import (
	"path/filepath"
	"testing"

	"tome/models"
)

// new memory db for testing
func newTestStore(t *testing.T) *Store {
	t.Helper()
	s, err := Open(":memory:")
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	t.Cleanup(func() { s.Close() })
	return s
}

// test new named level 1 character, dex 14, wis 16, wis spellcaster
func newTestCharacter(t *testing.T, name string) models.Character {
	t.Helper()
	c := models.NewCharacter()
	c.Base.SetName(name)
	c.SetScores(models.CharScores{Str: 10, Dex: 14, Con: 10, Int: 10, Wis: 16, Cha: 10})
	if err := c.SetSpellcastingAbility(models.Wis); err != nil {
		t.Fatalf("setup: SetSpellcastingAbility: %v", err)
	}
	if err := c.SetSkillLevel("Stealth", 1); err != nil {
		t.Fatalf("setup: SetSkillLevel: %v", err)
	}
	return c
}

func TestCreateAndLoadCharacter(t *testing.T) {
	s := newTestStore(t)
	want := newTestCharacter(t, "Bob")

	id, err := s.CreateCharacter(want)
	if err != nil {
		t.Fatalf("CreateCharacter: %v", err)
	}
	if id == 0 {
		t.Fatal("CreateCharacter returned id 0, want a real row id")
	}

	got, err := s.LoadCharacter(id)
	if err != nil {
		t.Fatalf("LoadCharacter(%d): %v", id, err)
	}

	// make sure row id takes precedent over json id
	want.ID = id
	if got != want {
		t.Errorf("loaded character differs from created one:\n got  %+v\n want %+v", got, want)
	}
}

func TestCreateCharacterStoresText(t *testing.T) {
	// make sure json_extract gets text and not []byte
	s := newTestStore(t)
	id, err := s.CreateCharacter(newTestCharacter(t, "Bob"))
	if err != nil {
		t.Fatalf("CreateCharacter: %v", err)
	}

	var typ string
	if err := s.db.QueryRow(`SELECT typeof(character) FROM characters WHERE id = ?`, id).Scan(&typ); err != nil {
		t.Fatalf("query column type: %v", err)
	}
	if typ != "text" {
		t.Errorf("character column stored as %q, want \"text\"", typ)
	}
}

func TestLoadCharacterRecalculates(t *testing.T) {
	// make sure derived stats are ignored from stored json and decoder recalculates properly
	s := newTestStore(t)
	c := newTestCharacter(t, "Bob")
	c.Skills.Stealth.Bonus = 99
	c.Spellcasting.SaveDC = 99

	id, err := s.CreateCharacter(c)
	if err != nil {
		t.Fatalf("CreateCharacter: %v", err)
	}
	got, err := s.LoadCharacter(id)
	if err != nil {
		t.Fatalf("LoadCharacter(%d): %v", id, err)
	}

	if got.Skills.Stealth.Bonus != 4 { // dex +2, PB 2
		t.Errorf("Stealth Bonus = %d, want 4", got.Skills.Stealth.Bonus)
	}
	if got.Spellcasting.SaveDC != 13 { // 8 + PB 2 + wis +3
		t.Errorf("SaveDC = %d, want 13", got.Spellcasting.SaveDC)
	}
}

func TestLoadCharacterNotFound(t *testing.T) {
	s := newTestStore(t)
	if _, err := s.LoadCharacter(999); err == nil {
		t.Error("LoadCharacter(999) on an empty database: got nil error, want not found")
	}
}

func TestSaveCharacter(t *testing.T) {
	s := newTestStore(t)
	id, err := s.CreateCharacter(newTestCharacter(t, "Bob"))
	if err != nil {
		t.Fatalf("CreateCharacter: %v", err)
	}
	c, err := s.LoadCharacter(id)
	if err != nil {
		t.Fatalf("LoadCharacter(%d): %v", id, err)
	}

	c.Base.SetName("Robert")
	c.SetLevel(5)
	if err := s.SaveCharacter(c); err != nil {
		t.Fatalf("SaveCharacter: %v", err)
	}

	got, err := s.LoadCharacter(id)
	if err != nil {
		t.Fatalf("LoadCharacter(%d) after save: %v", id, err)
	}
	if got.Base.Name != "Robert" {
		t.Errorf("Name = %q, want %q", got.Base.Name, "Robert")
	}
	if got.Base.Level != 5 || got.Proficiency != 3 {
		t.Errorf("Level, Proficiency = %d, %d, want 5, 3", got.Base.Level, got.Proficiency)
	}
	if got.Spellcasting.SaveDC != 14 { // 8 + PB 3 + wis +3
		t.Errorf("SaveDC = %d, want 14", got.Spellcasting.SaveDC)
	}
}

func TestSaveCharacterNotFound(t *testing.T) {
	tests := []struct {
		name string
		id   int64
	}{
		{"id never created", 999},
		{"zero id (never saved)", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := newTestStore(t)
			c := newTestCharacter(t, "Bob")
			c.ID = tt.id
			if err := s.SaveCharacter(c); err == nil {
				t.Errorf("SaveCharacter with ID %d: got nil error, want not found", tt.id)
			}
		})
	}
}

func TestListCharacters(t *testing.T) {
	s := newTestStore(t)

	list, err := s.ListCharacters()
	if err != nil {
		t.Fatalf("ListCharacters on empty database: %v", err)
	}
	// make sure Vue gets [] not nil/null
	if list == nil || len(list) != 0 {
		t.Errorf("empty ListCharacters = %#v, want an empty non-nil slice", list)
	}

	var want []models.CharacterSummary
	for _, name := range []string{"Bob", "Alice", ""} {
		id, err := s.CreateCharacter(newTestCharacter(t, name))
		if err != nil {
			t.Fatalf("CreateCharacter(%q): %v", name, err)
		}
		want = append(want, models.CharacterSummary{ID: id, Name: name})
	}

	list, err = s.ListCharacters()
	if err != nil {
		t.Fatalf("ListCharacters: %v", err)
	}
	if len(list) != len(want) {
		t.Fatalf("ListCharacters returned %d entries, want %d: %+v", len(list), len(want), list)
	}
	for i := range want {
		if list[i] != want[i] {
			t.Errorf("entry %d = %+v, want %+v", i, list[i], want[i])
		}
	}
}

func TestOpenKeepsDataBetweenOpens(t *testing.T) {
	// make sure reopening doesn't create new table over existing table
	path := filepath.Join(t.TempDir(), "tome.db")

	s, err := Open(path)
	if err != nil {
		t.Fatalf("first Open: %v", err)
	}
	id, err := s.CreateCharacter(newTestCharacter(t, "Bob"))
	if err != nil {
		t.Fatalf("CreateCharacter: %v", err)
	}
	if err := s.Close(); err != nil {
		t.Fatalf("Close: %v", err)
	}

	reopened, err := Open(path)
	if err != nil {
		t.Fatalf("second Open: %v", err)
	}
	defer reopened.Close()

	got, err := reopened.LoadCharacter(id)
	if err != nil {
		t.Fatalf("LoadCharacter(%d) after reopening: %v", id, err)
	}
	if got.Base.Name != "Bob" {
		t.Errorf("Name after reopening = %q, want %q", got.Base.Name, "Bob")
	}
}
