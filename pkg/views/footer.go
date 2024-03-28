package views

import (
	"fmt"
	"strings"

	"github.com/JuHaNi654/website-audit-tool/pkg/utils"
	"github.com/charmbracelet/lipgloss"
)

var (
	footerStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()
)

func (m Model) footer() string {
	info := footerStyle.Render(
		fmt.Sprintf(
			"%3.f%%",
			m.viewport.ScrollPercent()*100,
		),
	)

	w := m.viewport.Width - lipgloss.Width(info)
	line := strings.Repeat("-", utils.Max(0, w))

	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}
