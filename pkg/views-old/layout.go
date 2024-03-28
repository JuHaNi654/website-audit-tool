package views

import "fmt"

var (
  dot    = ColorFg(" • ", "#ffffff")
  subtle = MakeFgStyle("#ffffff", "") 
)

func Layout(title string, content string, m TUIMain) string {
  view := "\n"
  if title != "" {
    view += fmt.Sprintf("%s\n\n", subtle(title))
  }

  view += fmt.Sprintf("%s\n", content) 
  view += m.Help.View(m.Keys)

  return view 
}
