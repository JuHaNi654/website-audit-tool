package views

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
)

type inputModel struct {
	url textinput.Model
	err error
}

func newInputModel() *inputModel {
	urlField := textinput.New()
	urlField.Placeholder = "https://website.com"
	urlField.Focus()
	urlField.CharLimit = 155
	urlField.Width = 50
	return &inputModel{
		url: urlField,
	}
}

func renderSiteInput(m Model) string {
	view := m.input.url.View()

	if m.input.err != nil {
		view += fmt.Sprintf("\n%s", m.input.err.Error())
	}

	// TODO: display errors
	return layout(
		"Website url",
		view,
		m,
	)
}
