package main

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
)

func (m model) View() tea.View {
	var s strings.Builder

	switch m.screen {
	case "choose":
		m.chooseView(&s)
	case "confirm":
		m.confirmView(&s)
	case "":
		m.emptyView(&s)
	}

	teaView := tea.NewView(s.String())
	teaView.AltScreen = m.altscreen
	return teaView
}

func (m model) chooseView(s *strings.Builder) {
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

		fmt.Fprintf(s, "%s [%s] %s\n", cursor, checked, option)
	}
}

func (m model) confirmView(s *strings.Builder) {
	s.WriteString("Proceed with installation? [Y/n]\n")
}

func (m model) emptyView(s *strings.Builder) {
	s.WriteString("Info: Press enter to exit.\n")
}
