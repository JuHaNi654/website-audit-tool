package views

import (
	"github.com/JuHaNi654/website-audit-tool/pkg/validation"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func handleKeyInput(msg tea.KeyMsg, m Model) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch {
	case key.Matches(msg, m.keys.Quit):
		return m, tea.Quit
	case key.Matches(msg, m.keys.Enter):
		if m.viewState == MenuView {
			i := m.menu.items.Index()
			m.action = mapView[i]
			m.viewState = InputView
			return m, nil
		}

		if m.viewState == InputView {
			m.input.err = validation.ValidateUrl(m.input.url.Value())
			if m.input.err == nil {
				m.state = loading
				m.viewState = m.action
				return m, scan(m.input.url.Value(), m.action)
			}
		}

	case key.Matches(msg, m.keys.Back):
		m.viewState = MenuView
	}

	if m.viewState == InputView {
		m.input.url, cmd = m.input.url.Update(msg)
		cmds = append(cmds, cmd)
	}

	if m.viewState == MenuView {
		m.menu.items, cmd = m.menu.items.Update(msg)
		cmds = append(cmds, cmd)
	}

	if m.viewState == LinkView {
		m.linkResults, cmd = m.linkResults.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}
