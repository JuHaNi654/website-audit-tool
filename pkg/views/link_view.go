package views

import (
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

var tableStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

func renderDocumentLinks(m TUIMain) string {
	columns := []table.Column{
		{Title: "Label", Width: 20},
		{Title: "Link", Width: 40},
		{Title: "Status", Width: 6},
	}
	rows := []table.Row{}

	for _, item := range m.Result.Links {
		rows = append(rows, table.Row{
			item.Label,
			item.Link,
			fmt.Sprintf("%d", item.StatusCode),
		})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(false),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)

	t.SetStyles(s)

	return Layout(
		"Site links",
		tableStyle.Render(t.View()),
		m,
	)
}
