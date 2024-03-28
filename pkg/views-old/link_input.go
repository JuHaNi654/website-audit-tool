package views

import (
	"github.com/charmbracelet/bubbles/textinput"
)

func renderLinkInput(m TUIMain) string {
  view := m.Url.View() + "\n\n"

  if m.Error != nil {
    msg := m.Error.Error()
    view += ColorFg(msg, "#FF0000") 
  }

  return Layout(
    "Set website url",
    view,
    m,
  )
}

func newLinkInputModel() (ti textinput.Model) {
  ti = textinput.New()
  ti.Placeholder = "https://website.com"
  ti.Focus()
  ti.CharLimit = 155
  ti.Width = 45
  return 
}
