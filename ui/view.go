package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/lipgloss"
)

const (
	bullet = "•"
	up     = "↑"
	down   = "↓"
)

const (
	primaryColor   = lipgloss.Color("#b24c63")
	secondaryColor = lipgloss.Color("#fffbff")
	textColor      = lipgloss.Color("#89bbfe")
)

var (
	headFootStyle = lipgloss.NewStyle().Bold(true).Background(primaryColor).Foreground(secondaryColor)
	textStyle     = lipgloss.NewStyle().Foreground(textColor)
	symbolStyle   = lipgloss.NewStyle().Bold(true).Foreground(primaryColor)
)

func (m Model) View() string {
	var s strings.Builder

	switch m.state {
	case MESSAGE_CHOICE:
		s.WriteString(filePickerView(m))

	case MESSAGE_SEND:
		s.WriteString(spinnerView(m))

	case MESSAGE_SENT:
		s.WriteString(responseVew(m))
	}

	return s.String()
}

func filePickerView(m Model) string {
	var s strings.Builder

	s.WriteString(fmt.Sprintf("\n%s\n\n", headFootStyle.Render("Available Messages:")))

	for i, f := range m.filepicker.files {
		cursor := " "
		if m.filepicker.cursor == i {
			cursor = "*"
		}

		c := symbolStyle.Render(cursor)
		v := textStyle.Width(30).Render(f)
		s.WriteString(fmt.Sprintf("%s %s\n", c, v))
	}

	s.WriteString("\n")
	footer := fmt.Sprintf("%s/k up %s %s/j down %s ctrl+c quit", up, bullet, down, bullet)
	s.WriteString(headFootStyle.Render(footer))

	return s.String()
}

func spinnerView(m Model) string {
	var s strings.Builder

	m.spinner.Style = symbolStyle
	m.spinner.Spinner = spinner.Points

	s.WriteString("\n\n")
	s.WriteString(fmt.Sprintf(" %s %s", m.spinner.View(), textStyle.Render("Sending Message to SQS")))
	s.WriteString("\n\n")

	return s.String()
}

func responseVew(_ Model) string {
	var s strings.Builder

	s.WriteString("\n")
	s.WriteString(textStyle.Render("Message sent successfully"))

	s.WriteString("\n\n")
  footer := fmt.Sprintf("ctrl+c to quit %s enter to return", bullet)
	s.WriteString(headFootStyle.Render(footer))

	return s.String()
}
