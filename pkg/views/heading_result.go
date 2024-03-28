package views

import (
	"fmt"
	"strings"

	"github.com/JuHaNi654/website-audit-tool/pkg/audit"
)

func renderDocumentHeadings(m Model) string {
	s := ""
	var printH func([]*audit.Heading, int)
	printH = func(items []*audit.Heading, size int) {
		for _, item := range items {
			indent := strings.Repeat("\u2586", size)
			s += fmt.Sprintf("%s ", indent)
			s += item.Tag
			s += " \u254d "
			s += item.Text
			s += " \n"
			printH(item.Children, (size + 1))
		}
	}

	printH(m.result.Headings, 1)
	return s
}
