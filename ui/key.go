package ui

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Up    key.Binding
	Down  key.Binding
	Enter key.Binding
}

var DefaultKeyMap = KeyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
	),

	Down: key.NewBinding(
		key.WithKeys("down", "j"),
	),

	Enter: key.NewBinding(
		key.WithKeys("enter"),
	),
}
