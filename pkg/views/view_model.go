package views

import (
	"github.com/JuHaNi654/website-audit-tool/pkg/constant"
)

type action struct {
	action   int8
	choice   int8
	selected int8
	results  bool
}

func setupAction() *action {
	return &action{
		action:   0,
		choice:   0,
		selected: 0,
		results:  false,
	}
}

type view struct {
	baseView         constant.ViewState
	resultsView      constant.ViewState
	linkViewSettings linkView
}

func initView() *view {
	return &view{
		baseView: constant.MenuView,
		linkViewSettings: newLinkView(),
	}
}
