package main

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	tea "charm.land/bubbletea/v2"
)

func (m model) View() tea.View {
	var s strings.Builder
	switch m.screen {
	case "output":
		for i := range m.selected {
			s.WriteString(m.options[i])
			s.WriteString("\n")
		}
	case "choose":
		s.WriteString("Which dotfiles to install?\n\n")
		for i, option := range m.options {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}

			checked := " "
			if _, ok := m.selected[i]; ok {
				checked = "x"
			}

			fmt.Fprintf(&s, "%s [%s] %s\n", cursor, checked, option)
		}
		s.WriteString("\nPress q to quit.\n")
	}
	teaView := tea.NewView(s.String())
	teaView.AltScreen = true
	return teaView
}
