package styles

import "github.com/charmbracelet/lipgloss"

var (
	// Base styles
	Title = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		Padding(0, 1)

	SubTitle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#7D56F4"))

	// Layout
	AppContainer = lipgloss.NewStyle().
			Padding(1, 2)
)
