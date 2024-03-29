package ui

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/leberjs/pebble/internal/config"
)

const (
	MESSAGE_CHOICE = iota
	MESSAGE_SEND
	MESSAGE_SENT
)

type Model struct {
	config     config.Config
	filepicker filepicker
	spinner    spinner.Model
	state      int
}

type filepicker struct {
	cursor       int
	files        []string
	selectedFile int
}

func NewModel(c *config.Config, f []string) Model {
	s := spinner.New()

	m := Model{
		config:     *c,
		filepicker: filepicker{files: f},
		spinner:    s,
		state:      MESSAGE_CHOICE,
	}

	return m
}
