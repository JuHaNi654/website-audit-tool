package main

import (
	"github.com/JuHaNi654/web-check/pkg/views"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
  init := views.TUIMain{
    Action: 0,
    Choice: 0,
    Selected: 0,
    Url: views.NewLinkInputModel(),
  }

  p := tea.NewProgram(init, tea.WithAltScreen())
  if err := p.Start(); err != nil {
    fmt.Println("Could not start program:", err)
    os.Exit(1)
  }
}

