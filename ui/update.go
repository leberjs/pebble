package ui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case sendFileSqsMsg:
		return m, sendSqsMessageCmd(m)

	case sqsMessageSentMsg:
		m.state = MESSAGE_SENT
		// return m, tea.Quit

	case tea.KeyMsg:
		switch {
		case msg.Type == tea.KeyCtrlC:
			return m, tea.Quit

		case key.Matches(msg, DefaultKeyMap.Up):
			if m.filepicker.cursor > 0 {
				m.filepicker.cursor--
			}

		case key.Matches(msg, DefaultKeyMap.Down):
			if m.filepicker.cursor < len(m.filepicker.files)-1 {
				m.filepicker.cursor++
			}

		case key.Matches(msg, DefaultKeyMap.Enter):
			m.filepicker.selectedFile = m.filepicker.cursor
			m.state = MESSAGE_SEND

			cmd = sendSqsMessageCmd(m)
			cmds = append(cmds, cmd)
			cmds = append(cmds, m.spinner.Tick)
		}

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	return m, tea.Batch(cmds...)
}
