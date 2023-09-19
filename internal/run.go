package internal

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func Run() {
    var m Model
    if len(os.Args) > 1 {
        m = NewModel(os.Args[1])
    } else {
        // TODO: check for stored profile var in config
        m = NewModel("leberjs")
    }

    // p := tea.NewProgram(Model{}, tea.WithAltScreen())
    p := tea.NewProgram(m, tea.WithAltScreen())
    if _, err := p.Run(); err != nil {
        log.Fatal(err)
    }
}
