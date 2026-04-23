package main

import (
	tea "charm.land/bubbletea/v2"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	msgKeyPress, ok := msg.(tea.KeyPressMsg)
	if !ok {
		return m, nil
	}
	msgStr := msgKeyPress.String()

	var teaCmd tea.Cmd = nil
	switch msgStr {
	case "ctrl+c":
		return m, tea.Quit
	default:
		switch m.screen {
		case "select":
			teaCmd = m.selectUpdate(msgStr)

		case "confirm":
			teaCmd = m.confirmUpdate(msgStr)
		}
	}

	return m, teaCmd
}

func (m *model) selectUpdate(key string) tea.Cmd {
	switch key {
	case "up", "k":
		if m.selectCursor > 0 {
			m.selectCursor--
		}

	case "down", "j":
		if m.selectCursor < len(groups)-1 {
			m.selectCursor++
		}

	case "a":
		toggleOn := false
		if _, on := m.selected[0]; !on {
			toggleOn = true
		}
		if toggleOn {
			for k := range groups {
				m.selected[k] = struct{}{}
			}
		} else {
			for k := range groups {
				delete(m.selected, k)
			}
		}

	case "space":
		_, ok := m.selected[m.selectCursor]
		if ok {
			delete(m.selected, m.selectCursor)
		} else {
			m.selected[m.selectCursor] = struct{}{}
		}

	case "enter":
		m.setScreen("confirm")

	case "q":
		return tea.Quit
	}

	return nil
}

func (m *model) confirmUpdate(key string) tea.Cmd {
	switch key {
	case "y", "Y":
		m.confirmed = true
		return tea.Quit
	case "n", "N":
		m.setScreen("select")
	}

	return nil
}
