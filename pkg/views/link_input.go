package views

import "github.com/charmbracelet/bubbles/textinput"

func RenderLinkInput(m TUIMain) string {
  return Layout(
    "Set website url",
    m.Url.View(),
    m,
  )
}

func NewLinkInputModel() (ti textinput.Model) {
  ti = textinput.New()
  ti.Placeholder = "https://website.com"
  ti.Focus()
  ti.CharLimit = 155
  ti.Width = 45
  return 
}
