package internal

import tea "github.com/charmbracelet/bubbletea"

type Model struct {
	config AppConfig
}

type AppConfig struct {
	path    string
	profile string
}

func NewModel(profile string) Model {
	m := Model{
		config: AppConfig{profile: profile},
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m Model) View() string {
	return m.config.profile
}
