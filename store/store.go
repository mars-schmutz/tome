package store

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"tome/models"

	_ "modernc.org/sqlite"
)

type Store struct {
	db *sql.DB
}

func Open(path string) (*Store, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("connect to database: %w", err)
	}
	db.SetMaxOpenConns(1)

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS characters (
		id INTEGER PRIMARY KEY,
		character TEXT NOT NULL
		)`)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("create characters table: %w", err)
	}

	return &Store{db: db}, nil
}

func (s *Store) Close() error {
	return s.db.Close()
}

func (s *Store) CreateCharacter(c models.Character) (int64, error) {
	data, err := json.Marshal(c)
	if err != nil {
		return 0, fmt.Errorf("encode character: %w", err)
	}

	res, err := s.db.Exec(`INSERT INTO characters (character) VALUES (?)`, string(data))
	if err != nil {
		return 0, fmt.Errorf("insert character: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get new character id: %w", err)
	}
	return id, nil
}

func (s *Store) LoadCharacter(id int64) (models.Character, error) {
	var data []byte
	err := s.db.QueryRow(`SELECT character FROM characters WHERE id = ?`, id).Scan(&data)
	if errors.Is(err, sql.ErrNoRows) {
		return models.Character{}, fmt.Errorf("character %d not found", id)
	}
	if err != nil {
		return models.Character{}, fmt.Errorf("load character %d: %w", id, err)
	}

	c, err := models.DecodeCharacter(data)
	if err != nil {
		return models.Character{}, fmt.Errorf("load character %d: %w", id, err)
	}
	c.ID = id
	return c, nil
}

func (s *Store) SaveCharacter(c models.Character) error {
	data, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("encode character: %w", err)
	}

	res, err := s.db.Exec(`UPDATE characters SET character = ? WHERE id = ?`, string(data), c.ID)
	if err != nil {
		return fmt.Errorf("save character %d: %w", c.ID, err)
	}

	n, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("save character %d: %w", c.ID, err)
	}
	if n == 0 {
		return fmt.Errorf("save character %d: not found", c.ID)
	}
	return nil
}

func (s *Store) ListCharacters() ([]models.CharacterSummary, error) {
	rows, err := s.db.Query(`SELECT id, json_extract(character, '$.Base.Name') FROM characters ORDER BY id`)
	if err != nil {
		return nil, fmt.Errorf("list characters: %w", err)
	}
	defer rows.Close()

	list := []models.CharacterSummary{}
	for rows.Next() {
		var cs models.CharacterSummary
		if err := rows.Scan(&cs.ID, &cs.Name); err != nil {
			return nil, fmt.Errorf("list characters: %w", err)
		}
		list = append(list, cs)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("list characters: %w", err)
	}

	return list, nil
}
