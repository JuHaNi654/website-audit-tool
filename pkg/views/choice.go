package views

import (
	audit "github.com/JuHaNi654/website-audit-tool/pkg/audit"
	vs "github.com/JuHaNi654/website-audit-tool/pkg/constant"
	tea "github.com/charmbracelet/bubbletea"
)

func SelectChoiceSettings(msg tea.Msg, m TUIMain) (tea.Model, tea.Cmd) {
	if m.action.results {
		return defaultViewChoices(msg, &m)
	}

	if m.view.baseView == vs.InputView {
		return inputViewChoices(msg, &m)
	}

	return menuChoices(msg, &m)
}

func defaultViewChoices(msg tea.Msg, m *TUIMain) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab", "enter":
			m.action.results = false
			m.view.baseView = vs.MenuView
			return m, nil
		}
	}

	if m.view.resultsView == vs.LinkView {
		m.view.linkViewSettings.table, cmd = m.view.linkViewSettings.table.Update(msg)
		return m, cmd
	}

	return m, nil
}

func inputViewChoices(msg tea.Msg, m *TUIMain) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			m.view.baseView = vs.MenuView
			return m, nil
		case "enter":
			callScan(m)
			if m.Error != nil {
				return m, nil
			}
			m.action.results = true
			return m, nil
		}
	}
	m.Url, cmd = m.Url.Update(msg)
	return m, cmd
}

func callScan(m *TUIMain) {
	if m.view.resultsView == vs.HeadingView {
		m.Result, m.Error = audit.RunAudit(m.Url.Value(), vs.ScanHeading)
		return
	} else if m.view.resultsView == vs.LinkView {
		m.Result, m.Error = audit.RunAudit(m.Url.Value(), vs.ScanLinks)
	}
}

func menuChoices(msg tea.Msg, m *TUIMain) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			i := m.menu.list.Index()
			m.view.baseView = vs.InputView
			m.view.resultsView = vs.ViewState(i)
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.menu.list, cmd = m.menu.list.Update(msg)
	return m, cmd
}
