package main

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	cursor   int
	options  []string
	selected map[int]struct{}
}

func initialModel() model {
	m := model{
		cursor:   0,
		options:  []string{"git", "keyd", "zsh"},
		selected: make(map[int]struct{}),
	}
	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() tea.View {
	var s strings.Builder
	s.WriteString("What should we buy at the market?\n\n")
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

	return tea.NewView(s.String())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}
		case "space":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	return m, nil
}
