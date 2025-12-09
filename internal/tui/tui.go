package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/thelamedev/mattertui/internal/tui/model"
)

func Start() {
	p := tea.NewProgram(model.NewModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Display error: %v", err)
		os.Exit(1)
	}
}
