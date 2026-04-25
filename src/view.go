package main

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/fatih/color"
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
	if m.filterMode {
		s.WriteString(color.MagentaString("Filter: "))
	} else {
		s.WriteString("Filter: ")
	}
	s.WriteString(m.filterText + "\n")

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

		line := fmt.Sprintf("%s [%s] %s\n", cursor, checked, g.name)

		if m.selectCursor == i && !m.filterMode {
			line = color.MagentaString(line)
		}

		s.WriteString(line)
	}
}

func (m model) confirmView(s *strings.Builder) {
	s.WriteString("Proceed with installation? [Y/n]\n")
}
