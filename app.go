package main

import (
	"fmt"
	"os"

	"github.com/JuHaNi654/website-audit-tool/pkg/views"
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)


func main() {
  init := views.TUIMain{
    Action: 0,
    Choice: 0,
    Selected: 0,
    MenuList: views.NewMainMenu(),
    Url: views.NewLinkInputModel(),
    Keys: views.ControlKeys(),
    Help: help.New(),
  }
  
  p := tea.NewProgram(init, tea.WithAltScreen())
  if err := p.Start(); err != nil {
    fmt.Println("Could not start program:", err)
    os.Exit(1)
  }

}

