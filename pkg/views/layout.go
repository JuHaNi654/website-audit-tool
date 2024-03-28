package views

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var layoutStyle = lipgloss.NewStyle().Padding(0, 1)

func layout(title, content string, m Model) string {
	view := ""
	if title != "" {
		view += fmt.Sprintf("%s\n", title)
	}

	view += fmt.Sprintf("%s\n\n", content)
	view += m.help.View(m.keys)

	return layoutStyle.Render(view)
}
