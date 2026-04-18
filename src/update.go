package main

import (
	tea "charm.land/bubbletea/v2"
)

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
		case "a":
			toggleOn := false
			if _, on := m.selected[0]; !on {
				toggleOn = true
			}
			if toggleOn {
				for k := range m.options {
					m.selected[k] = struct{}{}
				}
			} else {
				for k := range m.options {
					delete(m.selected, k)
				}
			}
		case "space":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		case "tab":
			switch m.screen {
			case "output":
				m.screen = "choose"
			case "choose":
				m.screen = "output"
			}
		}
	}
	return m, nil
}
