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
			indentSpace := strings.Repeat("\u2586", indent)
			s += fmt.Sprintf("%s ", indentSpace)
			s += ColorFg(item.Tag, "#FAFAFA")
			s += " \u254d "
			s += ColorFg(item.Text, "#FAFAFA")
			s += " \n"
			printH(item.Children, (indent + 1))
		}
	}

	printH(m.Result.Headings, 1)
	return Layout(
		"Site headings result",
		s,
		m,
	)
}
