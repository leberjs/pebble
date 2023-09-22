package tui

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/leberjs/pebble/internal/config"
)

func Run(c *config.Config, f []string) {
	m := NewModel(c, f)

	p := tea.NewProgram(m)
	// p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
