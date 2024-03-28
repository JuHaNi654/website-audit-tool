package views

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	itemStyle         = lipgloss.NewStyle().Padding(0, 1)
	selectedItemStyle = lipgloss.NewStyle().Padding(0, 1).Foreground(lipgloss.Color("170"))
)

type menu struct {
	items list.Model
}

func newMenu() *menu {
	return &menu{
		items: newList(),
	}
}

type menuKeyMap struct {
	toggleSelect key.Binding
	toggleQuit   key.Binding
}

func newKeyMap() *menuKeyMap {
	return &menuKeyMap{
		toggleSelect: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "Select"),
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
func (d menuItemDelegate) Render(
	w io.Writer,
	m list.Model,
	index int,
	listItem list.Item,
) {
	i, ok := listItem.(menuItem)
	if !ok {
		return
	}

	str := fmt.Sprintf("%s", i)
	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

func newList() list.Model {
	listKeys := newKeyMap()

	items := []list.Item{
		menuItem("List page headings"),
		menuItem("List page links"),
	}

	l := list.New(items, menuItemDelegate{}, 20, 30)
	l.SetShowTitle(false)
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
