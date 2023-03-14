package views

import (
	audit "github.com/JuHaNi654/website-audit-tool/pkg/audit"
	vs "github.com/JuHaNi654/website-audit-tool/pkg/constant"
	tea "github.com/charmbracelet/bubbletea"
)

func SelectChoiceSettings(msg tea.Msg, m TUIMain) (tea.Model, tea.Cmd) {
	if m.ShowResults {
		return defaultViewChoices(msg, &m)
	}

	if m.BaseView == vs.InputView {
		return inputViewChoices(msg, &m)
	}

	return menuChoices(msg, &m)
}

func defaultViewChoices(msg tea.Msg, m *TUIMain) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab", "enter":
			m.ShowResults = false
			m.BaseView = vs.MenuView
			return m, nil
		}
	}

	return m, nil
}

func inputViewChoices(msg tea.Msg, m *TUIMain) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			m.BaseView = vs.MenuView
			return m, nil
		case "enter":
			callScan(m)
			if m.Error != nil {
				return m, nil
			}
			m.ShowResults = true
			return m, nil
		}
	}
	m.Url, cmd = m.Url.Update(msg)
	return m, cmd
}

func callScan(m *TUIMain) {
	if m.ViewState == vs.HeadingView {
		m.Result, m.Error = audit.RunAudit(m.Url.Value(), vs.ScanHeading)
		return
	} else if m.ViewState == vs.LinkView {
		m.Result, m.Error = audit.RunAudit(m.Url.Value(), vs.ScanLinks)
	}
}

func menuChoices(msg tea.Msg, m *TUIMain) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			i := m.MenuList.Index()
			m.BaseView = vs.InputView
			m.ViewState = vs.ViewState(i)
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.MenuList, cmd = m.MenuList.Update(msg)
	return m, cmd
}
