package views

import (
	"fmt"

	audit "github.com/JuHaNi654/website-audit-tool/pkg/audit"
	vs "github.com/JuHaNi654/website-audit-tool/pkg/constant"
	"github.com/charmbracelet/bubbles/help"
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
	action *action
	menu   *menu
	view   *view
	Url    textinput.Model
	Result *audit.HtmlDocumentAudit
	Keys   KeyMap
	Help   help.Model
	Error  error
}

func NewInstance() TUIMain {
	return TUIMain{
		action: setupAction(),
		menu:   newMenu(),
		view:   initView(),
		Url:    newLinkInputModel(),
		Keys:   controlKeys(),
		Help:   help.New(),
	}
}

func (m TUIMain) Init() tea.Cmd {
	return nil
}

func (m TUIMain) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.menu.list.SetWidth(msg.Width)
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

	if m.action.results {
		switch m.view.resultsView {
		case vs.HeadingView:
			s += RenderDocumentHeadings(m)
		case vs.LinkView:
			s += renderDocumentLinks(m)
		}
	} else {
		switch m.view.baseView {
		case vs.MenuView:
			s += m.menu.list.View()
		case vs.InputView:
			s += renderLinkInput(m)
		}
	}

	return indent.String("\n"+s+"\n\n", 2)
}
