package views

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/lipgloss"
)

var (
	spinnerSpacing = lipgloss.NewStyle().Padding(0, 1)
	spinnerStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
)

func newSpinner() spinner.Model {
	s := spinner.New()
	s.Style = spinnerStyle
	s.Spinner = spinner.Dot
	return s
}

func spinnerView(m Model) string {
	return spinnerSpacing.Render(fmt.Sprintf(
		"%s %s",
		m.spinner.View(),
		"Loading ...",
	))
}
