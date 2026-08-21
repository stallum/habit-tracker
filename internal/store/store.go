package store

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type Store struct {
	db *sql.DB
}

func Open(path string) (*Store, error) {
	db, err := sql.Open("sqlite", path+"?_pragma=foregein_keys(1)")
	if err != nil {
		return nil, fmt.Errorf("abrir banco: %w", err)
	}

	if err := migrate(db); err != nil {
		db.Close()
		return nil, fmt.Errorf("migrar banco: %w", err)
	}

	return &Store{db: db}, nil
}

func (s *Store) Close() error {
	return s.db.Close()
}

func migrate(db *sql.DB) error {
	const schema = `
	CREATE TABLE IF NOT EXISTS habits (
		id			INTEGER PRIMARY KEY AUTOINCREMENT,
		name		TEXT NOT NULL,
		created_at	DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS completions (
		id       INTEGER PRIMARY KEY AUTOINCREMENT,
		habit_id INTEGER NOT NULL REFERENCES habits(id) ON DELETE CASCADE,
		date     DATETIME,
		UNIQUE(habit_id, date)
	);

	CREATE TABLE IF NOT EXISTS settings (
		key   TEXT PRIMARY KEY,
		value TEXT NOT NULL
	);
	`

	_, err := db.Exec(schema)
	return err
}