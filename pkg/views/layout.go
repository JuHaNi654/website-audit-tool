package views

import "fmt"

var (
  dot    = ColorFg(" • ", "#ffffff")
  subtle = MakeFgStyle("#ffffff", "") 
)

func Layout(title string, content string, m TUIMain) string {
  view := "\n"
  if title != "" {
    view += fmt.Sprintf("%s\n\n", title)
  }

  view += content 
  view += controls(m)

  return view 
}


func controls(m TUIMain) string {
  s := "\n\n"
  if m.RenderView == InputView {
    return s + subtle("esc: quit") + dot + subtle("ctrl+b: cancel") + dot + subtle("enter: continue")
  }

  if m.RenderView == HeadingView {
    return s + subtle("0: go back")
  }

  return s + subtle("esc: quit")
}
