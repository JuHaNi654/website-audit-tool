package views

import (
	"fmt"

	"github.com/JuHaNi654/website-audit-tool/pkg/audit"
	"github.com/JuHaNi654/website-audit-tool/pkg/debug"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type linkItem struct {
	title string
}

func newLinkItem(node *audit.LinkNode) linkItem {
	return linkItem{
		title: node.Label,
	}
}

func (i linkItem) FilterValue() string { return "" }
func (i linkItem) Title() string       { return i.title }

func initLinkList() list.Model {
	l := list.New(nil, list.NewDefaultDelegate(), 0, 0)
	l.SetShowTitle(false)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.DisableQuitKeybindings()

	return l
}

func newLinkItemList(m *Model) {
	items := make([]list.Item, len(m.result.Links))
	for i, link := range m.result.Links {
		items[i] = newLinkItem(link)
	}

	m.linkResults.SetItems(items)
}

func newItemDelegate() list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
		debug.Debugger(fmt.Sprintf("%s", m.SelectedItem()))
		return nil
	}

	return d
}
