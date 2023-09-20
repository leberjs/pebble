package tui

import "github.com/leberjs/pebble/internal/config"

type Model struct {
	config config.Config
}

func NewModel(c *config.Config) Model {
    m := Model{config: *c}

	return m
}
