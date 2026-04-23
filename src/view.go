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
	s.WriteString("Which dotfiles to install?\n\n")
	for i, option := range groups {
		cursor := " "
		if m.selectCursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		fmt.Fprintf(s, "%s [%s] %s\n", cursor, checked, option.name)
	}
}

func (m model) confirmView(s *strings.Builder) {
	s.WriteString("Proceed with installation? [Y/n]\n")
}
