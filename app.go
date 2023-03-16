package main

import (
	"fmt"
	"os"

	"github.com/JuHaNi654/website-audit-tool/pkg/views"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	init := views.NewInstance()

	p := tea.NewProgram(init, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Println("Could not start program:", err)
		os.Exit(1)
	}
}
