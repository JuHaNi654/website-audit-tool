package main

import (
	"fmt"
	"os"

	"github.com/JuHaNi654/website-audit-tool/pkg/constant"
	"github.com/JuHaNi654/website-audit-tool/pkg/views"
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	init := views.TUIMain{
		Action:      0,
		Choice:      0,
		Selected:    0,
		MenuList:    views.NewMainMenu(),
		Url:         views.NewLinkInputModel(),
		Keys:        views.ControlKeys(),
		Help:        help.New(),
		ShowResults: false,
		BaseView:    constant.MenuView,
	}

	p := tea.NewProgram(init, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Println("Could not start program:", err)
		os.Exit(1)
	}

	// audit.RunAudit("http://localhost:3000/test.html", 1)
}
