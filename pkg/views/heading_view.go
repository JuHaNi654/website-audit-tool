package views

import (
	"fmt"
	"strings"

	audit "github.com/JuHaNi654/website-audit-tool/pkg/audit"
)

func RenderDocumentHeadings(m TUIMain) string {
  s := ""
  var printH func([]*audit.Heading, int) 
  printH = func(items []*audit.Heading, indent int) {
    for _, item := range items {
      indentSpace := strings.Repeat(" ", indent)
      s += fmt.Sprintf("%s- %s: %s\n", indentSpace, item.Tag, item.Text)
      printH(item.Children, (indent + 1))
    }
  }

  printH(m.Result.Headings, 0)
  return Layout(
    "Site headings result",
    s,
    m,
  )
}
