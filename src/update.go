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
		case "choose":
			teaCmd = m.chooseUpdate(msgStr)

		case "confirm":
			teaCmd = m.confirmUpdate(msgStr)

		case "":
			teaCmd = m.emptyUpdate(msgStr)
		}
	}

	return m, teaCmd
}

func (m *model) chooseUpdate(key string) tea.Cmd {
	switch key {
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
		m.setScreen("")
	case "n", "N":
		m.setScreen("choose")
	}

	return nil
}

func (m *model) emptyUpdate(key string) tea.Cmd {
	switch key {
	case "enter":
		return tea.Quit
	}

	return nil
}
