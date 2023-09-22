package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/leberjs/pebble/internal/aws"
	"github.com/leberjs/pebble/internal/syncer"
)

func sendSqsMessageCmd(m Model) tea.Cmd {
	return func() tea.Msg {
		fn := m.filepicker.files[m.filepicker.selectedFile]
		fc := syncer.GetFileContent(fn)

		aws.AwsConfig(m.config.ProfileName()).SendMessage(fc, m.config.QueueUrl())

		return sqsMessageSentMsg{}
	}
}
