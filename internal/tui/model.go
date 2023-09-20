package tui

import tea "github.com/charmbracelet/bubbletea"

type Model struct {
	config AppConfig
}

type AppConfig struct {
	path        string
	profileName string
	syncBucket  string
	queueUrl    string
}

func NewModel(profileName, syncBucket, queueUrl string) Model {
	m := Model{
		config: AppConfig{
			profileName: profileName,
			syncBucket:  syncBucket,
			queueUrl:    queueUrl,
		},
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
	return m.config.profileName
}
