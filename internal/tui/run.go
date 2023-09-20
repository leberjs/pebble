package tui

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/leberjs/pebble/internal/config"
)

type RunContext struct {
	profileName string
	syncBucket  string
	queueUrl    string
}

func Run(c *config.Config) {
	m := NewModel(c)

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
