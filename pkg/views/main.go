package views

import (
	"fmt"
	"os"

	audit "github.com/JuHaNi654/website-audit-tool/pkg/audit"
	vs "github.com/JuHaNi654/website-audit-tool/pkg/constant"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/indent"
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
	Action      int8
	Choice      int8
	Selected    int8
	MenuList    list.Model
	ShowResults bool
	BaseView    vs.ViewState
	ViewState   vs.ViewState
	Url         textinput.Model
	Result      *audit.HtmlDocumentAudit
	Keys        KeyMap
	Help        help.Model
	Error       error
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
	return SelectChoiceSettings(msg, m)
}

func (m TUIMain) View() string {
	s := fmt.Sprintf("%s\n\n", titleStyle.Render("Website audit tool"))

	if m.ShowResults {
		switch m.ViewState {
		case vs.HeadingView:
			s += RenderDocumentHeadings(m)
		case vs.LinkView:
			s += renderDocumentLinks(m)
		}
	} else {
		switch m.BaseView {
		case vs.MenuView:
			s += m.MenuList.View()
		case vs.InputView:
			s += RenderLinkInput(m)
		}
	}

	return indent.String("\n"+s+"\n\n", 2)
}

func logErrors(e error) {
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
