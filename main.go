package main 

import (
	"fmt"
	"os"

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
	p := tea.NewProgram(model{message: "Olá! Aperte \"q\" para sair.\n"})
	if _, err := p.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "erro:", err)
		os.Exit(1)
	}
}