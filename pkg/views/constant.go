package views

import "github.com/JuHaNi654/website-audit-tool/pkg/audit"

type ViewState uint8
type AppState int

const (
	initializing AppState = iota
	loading
	ready
	other // error
)

const (
	MenuView ViewState = iota
	HeadingView
	InputView
	LinkView
	ErrorView
)

var mapAction = map[ViewState]audit.ScanAction{
	HeadingView: audit.ScanHeading,
	LinkView:    audit.ScanLinks,
}

var mapView = map[int]ViewState{
	0: HeadingView,
	1: LinkView,
}
