package views

import "github.com/muesli/termenv"

var (
	term = termenv.EnvColorProfile()
)

func ColorFg(val, color string) string {
	return termenv.String(val).Foreground(term.Color(color)).String()
}

func MakeFgStyle(fg, bg string) func(string) string {
	return termenv.Style{}.Foreground(term.Color(fg)).Background(term.Color(bg)).Styled
}
