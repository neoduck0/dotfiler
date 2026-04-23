package main

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
)

func (m model) View() tea.View {
	var s strings.Builder

	switch m.screen {
	case "select":
		m.selectView(&s)
	case "confirm":
		m.confirmView(&s)
	}

	teaView := tea.NewView(s.String())
	teaView.AltScreen = m.altscreen
	return teaView
}

func (m model) selectView(s *strings.Builder) {
	s.WriteString("Filter: " + m.filterText + "\n")

	for i, g := range m.filterList {
		cursor := " "
		if !m.filterMode {
			if m.selectCursor == i {
				cursor = ">"
			}
		}

		checked := " "
		if g.selected {
			checked = "x"
		}

		fmt.Fprintf(s, "%s [%s] %s\n", cursor, checked, g.name)
	}
}

func (m model) confirmView(s *strings.Builder) {
	s.WriteString("Proceed with installation? [Y/n]\n")
}
