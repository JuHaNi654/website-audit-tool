package views

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	defaultWidth  = 20
	defaultHeight = 14
)

var (
	itemStyle         = lipgloss.NewStyle()
	selectedItemStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("170"))
)

type menu struct {
	list list.Model
}

func newMenu() *menu {
	return &menu{
		list: newList(),
	}
}

type listKeyMap struct {
	toggleSelect key.Binding
	toggleQuit   key.Binding
}

func newListKeyMap() *listKeyMap {
	return &listKeyMap{
		toggleSelect: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "Select action"),
		),
		toggleQuit: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "Quit"),
		),
	}
}

type menuItem string

func (i menuItem) FilterValue() string { return "" }

type menuItemDelegate struct{}

func (d menuItemDelegate) Height() int                               { return 1 }
func (d menuItemDelegate) Spacing() int                              { return 0 }
func (d menuItemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d menuItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(menuItem)
	if !ok {
		return
	}

	str := fmt.Sprintf("%s", i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s string) string {
			return selectedItemStyle.Render("> " + s)
		}
	}

	fmt.Fprint(w, fn(str))
}

func newList() list.Model {
	listKeys := newListKeyMap()

	items := []list.Item{
		menuItem("Print page header layout"),
		menuItem("Crawl page links"),
	}

	l := list.New(items, menuItemDelegate{}, defaultWidth, defaultHeight)
	l.Title = "Choose action"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.DisableQuitKeybindings()
	l.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			listKeys.toggleSelect,
			listKeys.toggleQuit,
		}
	}
	return l
}
