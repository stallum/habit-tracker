package main 

import (
	"fmt"
	"log"

	// "github.com/stallum/habit-tracker/internal/store"
	"habit-tracker/internal/store"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	message string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "ctrl+c": 
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() tea.View {
	return tea.NewView(m.message)
}

func main() {
	s, err := store.Open("habits.db")
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	h, err := s.CreateHabit("Beber Água")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Criado: %+v\n", h)

	habits, err := s.ListHabits()
	if err != nil {
		log.Fatal(err)
	}

	for _, h := range habits {
		fmt.Printf(
			"- [%d] %s (Criado em %s)\n",
			h.ID,
			h.Name,
			h.CreatedAt.Format("02/01/2006"),
		)
	}
}