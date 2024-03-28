package views

import (
	"github.com/charmbracelet/bubbles/viewport"
)

func (m *Model) handleViewport(w, h int) {
	//	headerHeight := lipgloss.Height(m.header())
	//	footerHeight := lipgloss.Height(m.footer())
	//	verticalMargin := headerHeight + footerHeight
	verticalMargin := 0
	headerHeight := 0

	m.menu.items.SetSize(w, h-verticalMargin)
	m.linkResults.SetSize(w, h-verticalMargin)

	if m.result != nil {
		m.viewport = viewport.New(w, h-verticalMargin)
		m.viewport.YPosition = headerHeight
		m.viewport.HighPerformanceRendering = useHighPerformanceRenderer
		m.viewport.SetContent(renderDocumentHeadings(*m))

		m.viewport.YPosition = headerHeight + 1
	}

	if m.state == initializing {
		m.viewport = viewport.New(w, h-verticalMargin)
		m.state = ready
	} else {
		m.viewport.Width = w
		m.viewport.Height = h - verticalMargin
	}
}
