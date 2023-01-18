package views

import (
	"fmt"
	"os"

	audit "github.com/JuHaNi654/website-audit-tool/pkg/audit"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/indent"
)


type ViewState int64 

const (
  MenuView ViewState = iota
  InputView
  HeadingView 
)

var border = lipgloss.Border{
  Top: "*",
  Bottom:"*",
}

var titleStyle = lipgloss.NewStyle().
    BorderStyle(border).
    Bold(true).
    Foreground(lipgloss.Color("#FAFAFA"))


type TUIMain struct {
	Action   int8
	Choice   int8
	Selected int8
  RenderView ViewState
  Url textinput.Model 
  Result *audit.HtmlDocumentAudit
  Error error
}

func (m TUIMain) Init() tea.Cmd {
  return nil 
}

func (m TUIMain) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  f, err := tea.LogToFile("debug.log", "debug")
  if err != nil {
    fmt.Println("fatal: ", err)
    os.Exit(1)
  }

  defer f.Close()

  if m.Error != nil {
    f.Write([]byte(m.Error.Error()))
    f.Write([]byte("\n"))
  }

  if msg, ok := msg.(tea.KeyMsg); ok {
    k := msg.String()
    if k == "ctrl+c" || k == "esc" {
      return m, tea.Quit
    }
  }
 if m.RenderView == InputView {
    return inputViewChoices(msg, &m)
  } else if m.RenderView == HeadingView {
    return defaultViewChoices(msg, &m)
  }

  return menuChoices(msg, &m)
}

func (m TUIMain) View() string {
  s := fmt.Sprintf("%s\n\n", titleStyle.Render("Website audit tool"))

  switch m.RenderView {
  case MenuView:
    s += MainMenu(m)
  case HeadingView:
    s += RenderDocumentHeadings(m)
  case InputView:
    s += RenderLinkInput(m)
  }

  return indent.String("\n"+s+"\n\n", 2)
}


func defaultViewChoices(msg tea.Msg, m *TUIMain) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case tea.KeyMsg:
    switch msg.String() {
    case "0":
      m.RenderView = MenuView
    }
  }

  return m, nil
}

func inputViewChoices(msg tea.Msg, m *TUIMain) (tea.Model, tea.Cmd) {
  var cmd tea.Cmd
  var err error
  switch msg := msg.(type) {
  case tea.KeyMsg:
    switch msg.String() {
    case "ctrl+b":
      m.RenderView = MenuView
      return m, nil 
    case "enter":
      m.Result, err = audit.RunAudit(m.Url.Value())
      if err != nil {
        m.Error = err 
        return m, nil 
      }

      m.Error = nil
      m.RenderView = HeadingView
      return m, nil
    }
  }
  m.Url, cmd = m.Url.Update(msg)
  return m, cmd 
}

func menuChoices(msg tea.Msg, m *TUIMain) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case tea.KeyMsg:
    switch msg.String() {
    case "1":
      m.RenderView = InputView
    }
  }
  
  return m, nil 
}
