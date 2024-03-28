package views

import (
	"github.com/JuHaNi654/website-audit-tool/pkg/audit"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

const useHighPerformanceRenderer = false

var ()

type Model struct {
	action      ViewState
	viewState   ViewState
	viewport    viewport.Model
	menu        *menu
	input       *inputModel
	state       AppState // TODO:  change state->status
	result      *audit.HtmlDocumentAudit
	err         error
	linkResults list.Model

	spinner spinner.Model

	keys keyMap
	help help.Model
}

func NewInstance() Model {
	return Model{
		viewState:   MenuView,
		menu:        newMenu(),
		input:       newInputModel(),
		state:       initializing,
		keys:        controlKeys(),
		help:        help.New(),
		spinner:     newSpinner(),
		linkResults: initLinkList(),
	}
}

func (m Model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)
	switch msg := msg.(type) {
	case *audit.HtmlDocumentAudit:
		m.result = msg
		m.state = ready

		if m.viewState == HeadingView {
			// TODO Display message if no headings founds
			m.viewport.SetContent(renderDocumentHeadings(m))
		}

		if m.viewState == LinkView {
			// TODO Display message if no links found
			newLinkItemList(&m)
		}
	case error:
		m.err = msg
		m.state = ready
	case tea.KeyMsg:
		return handleKeyInput(msg, m)
	case tea.WindowSizeMsg:
		m.handleViewport(msg.Width, msg.Height)

		if useHighPerformanceRenderer {
			cmds = append(cmds, viewport.Sync(m.viewport))
		}
	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	s := ""
	//s := fmt.Sprintf("%s\n", m.header())

	if m.state == initializing {
		return "Initializing app ..."
	}

	if m.state == loading {
		s += spinnerView(m)
		return s
	}

	if m.viewState == HeadingView {
		s += m.viewport.View()
	} else if m.viewState == LinkView {
		s += m.linkResults.View()
	} else if m.viewState == MenuView {
		s += m.menu.items.View()
	} else if m.viewState == InputView {
		s += renderSiteInput(m)
	}

	//s += m.footer()
	return s
}
