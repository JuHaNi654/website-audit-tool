package views

import (
	"github.com/JuHaNi654/website-audit-tool/pkg/audit"
	tea "github.com/charmbracelet/bubbletea"
)

func scan(url string, vs ViewState) tea.Cmd {
	return func() tea.Msg {
		result, err := audit.RunAudit(
			url,
			mapAction[vs],
		)

		if err != nil {
			return err
		}

		return result
	}
}
