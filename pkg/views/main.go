package views

import (
	"fmt"
	"os"

	audit "github.com/JuHaNi654/website-audit-tool/pkg/audit"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/indent"
)


type ViewState int8 

const (
  MenuView ViewState = iota
  InputView
  HeadingView 
)

var border = lipgloss.Border{
		Top:         "\u2550",
		Bottom:      "\u2550",
		Left:        "\u2551",
		Right:       "\u2551",
		TopLeft:     "\u2554",
		TopRight:    "\u2557",
		BottomLeft:  "\u255a",
		BottomRight: "\u255d",
}

var titleStyle = lipgloss.NewStyle().
    BorderStyle(border).
    Bold(true).
    Foreground(lipgloss.Color("#FAFAFA"))


type TUIMain struct {
	Action   int8
	Choice   int8
	Selected int8
  MenuList list.Model
  RenderView ViewState
  Url textinput.Model 
  Result *audit.HtmlDocumentAudit
  Keys KeyMap
  Help help.Model
  Error error
}

func (m TUIMain) Init() tea.Cmd {
  return nil 
}

func (m TUIMain) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case tea.WindowSizeMsg:
    m.MenuList.SetWidth(msg.Width)
    return m, nil
  case tea.KeyMsg:
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
    s += m.MenuList.View()
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
    case "tab", "enter":
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
    case "tab":
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
    case "enter":
      i := m.MenuList.Index()
      m.RenderView = ViewState(i+1)
      return m, nil
    }
  }

  var cmd tea.Cmd
  m.MenuList, cmd = m.MenuList.Update(msg)
  return m, cmd 
}


func logErrors (e error) {
  f, err := tea.LogToFile("debug.log", "debug")
  if err != nil {
    fmt.Println("fatal: ", err)
    os.Exit(1)
  }

  defer f.Close()

  if e != nil {
    f.Write([]byte(e.Error()))
    f.Write([]byte("\n"))
  }

}
