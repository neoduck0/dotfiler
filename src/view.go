package main

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (m *model) View() tea.View {
	var s strings.Builder

	header := renderStyle(titleStyle, "Dotfiler")
	s.WriteString(header + "\n\n")

	switch m.currentView {
	case selectView:
		m.selectView(&s)
	case confirmView:
		m.confirmView(&s)
	}

	teaView := tea.NewView(s.String())
	teaView.AltScreen = true
	return teaView
}

func (m *model) selectView(s *strings.Builder) {
	if m.filterInput.Focused() {
		s.WriteString(renderStyle(selectedStyle, "Filter: "))
	} else {
		s.WriteString("Filter: ")
	}
	s.WriteString(m.filterInput.View() + "\n")

	for i, g := range m.filterList {
		cursor := " "
		if !m.filterInput.Focused() {
			if m.selectCursor == i {
				cursor = ">"
			}
		}

		checked := " "
		if g.selected {
			checked = "x"
		}

		line := fmt.Sprintf("%s [%s] %s\n", cursor, checked, g.name)

		if m.selectCursor == i && !m.filterInput.Focused() {
			line = renderStyle(selectedStyle, line)
		}

		s.WriteString(line)
	}
}

func (m *model) confirmView(s *strings.Builder) {
	s.WriteString("Proceed with installation? [Y/n]\n")
}

func renderStyle(style lipgloss.Style, s string) string {
	return strings.TrimSpace(style.Render(s))
}

var (
	titleStyle    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Green)
	selectedStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Magenta)
)
