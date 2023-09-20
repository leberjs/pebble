package tui

func (m Model) View() string {
	return m.config.ProfileName()
}
