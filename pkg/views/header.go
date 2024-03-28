package views

import (
	"strings"

	"github.com/JuHaNi654/website-audit-tool/pkg/utils"
	"github.com/charmbracelet/lipgloss"
)

var (
	headerStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()
)

func (m Model) header() string {
	title := headerStyle.Render("Website audit tool")
	w := m.viewport.Width - lipgloss.Width(title)
	line := strings.Repeat("-", utils.Max(0, w))

	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}
