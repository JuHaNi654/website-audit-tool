package views

import (
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

var tableStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type linkView struct {
	table table.Model
}

func newLinkView() linkView {
	return linkView{
		table: newLinkTable(),
	}
}

func newLinkTable() table.Model {
	columns := []table.Column{
		{Title: "Label", Width: 20},
		{Title: "Link", Width: 40},
		{Title: "Status", Width: 6},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithFocused(true),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)

	s.Selected = s.Selected.
		Foreground(lipgloss.Color("240")).
		Background(lipgloss.Color("15")).
		Bold(false)
	t.SetStyles(s)
	return t
}

func renderDocumentLinks(m TUIMain) string {
	rows := []table.Row{}

	for _, item := range m.Result.Links {
		rows = append(rows, table.Row{
			item.Label,
			item.Link,
			fmt.Sprintf("%d", item.StatusCode),
		})
	}

	m.view.linkViewSettings.table.SetRows(rows)
	return Layout(
		fmt.Sprintf("Total links %d", len(m.Result.Links)),
		tableStyle.Render(m.view.linkViewSettings.table.View()),
		m,
	)
}
