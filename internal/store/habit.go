package store

import (
	"fmt"
	"time"
)

type Habit struct {
	ID			int64
	Name 		string
	CreatedAt	time.Time
}

func (s *Store) CreateHabit(name string) (Habit, error) {
	res, err := s.db.Exec(`INSERT INTO habits (name) VALUES (?)`, name)
	if err != nil {
		return Habit{}, fmt.Errorf("Criar habito: %w", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return Habit{}, fmt.Errorf("criar hábito: %w", err)
	}

	return s.GetHabit(id)
}

func (s *Store) GetHabit(id int64) (Habit, error) {
	var h Habit
	// var createdAt string

	err := s.db.QueryRow(`SELECT id, name, created_at FROM habits WHERE id = ?`, id).Scan(&h.ID, &h.Name, &h.CreatedAt)

	if err != nil {
		return Habit{}, fmt.Errorf("buscar hábito: %w", err)
	}

	// h.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt)
	return h, nil
}

func (s *Store) ListHabits() ([]Habit, error) {
	rows, err := s.db.Query(`SELECT id, name, created_at FROM habits ORDER BY id`)
	if err != nil {
		return nil, fmt.Errorf("listar hábitos: %w", err)
	}

	defer rows.Close()

	var habits []Habit
	for rows.Next() {
		var h Habit
		// var createdAt string

		if err := rows.Scan(&h.ID, &h.Name, &h.CreatedAt); err != nil {
			return nil, fmt.Errorf("listar hábitos: %w", err)
		}

		// h.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", h.CreatedAt)
		habits = append(habits, h)
	}
	return habits, rows.Err()
}

func (s *Store) DeleteHabit(id int64) error {
	_, err := s.db.Exec(`DELETE FROM habits WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("deletar hábito: %w", err)
	}
	return nil
}

