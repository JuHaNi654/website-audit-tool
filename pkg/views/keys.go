package views

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Quit  key.Binding
	Enter key.Binding
	Back  key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Enter, k.Back, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding { return [][]key.Binding{} }

func controlKeys() keyMap {
	return keyMap{
		Quit: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "(Quit)"),
		),
		Enter: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "(Continue)"),
		),
		Back: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "(Go back)"),
		),
	}
}
