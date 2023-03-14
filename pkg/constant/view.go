package constant

type ViewState uint8 
type ScanAction uint8

const (
  MenuView ViewState = iota
  InputView
)

const (
  HeadingView ViewState = iota
  LinkView
)

const (
  ScanHeading ScanAction = iota 
  ScanLinks
)
