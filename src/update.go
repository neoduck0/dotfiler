package main

import (
	tea "charm.land/bubbletea/v2"
)

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		switch m.currentView {
		case selectView:
			teaCmd = m.selectUpdate(msgStr)

		case confirmView:
			teaCmd = m.confirmUpdate(msgStr)
		}
	}

	return m, teaCmd
}

func (m *model) selectUpdate(key string) tea.Cmd {
	switch key {
	case "tab":
		m.filterMode = !m.filterMode
		return nil

	case "/":
		m.filterMode = true
		return nil

	case "esc":
		m.filterMode = false
		return nil

	case "enter":
		if m.filterMode {
			m.filterMode = false
			return nil
		}
	}

	if m.filterMode {
		if key == "backspace" && len(m.filterText) >= 1 {
			m.filterText = m.filterText[0 : len(m.filterText)-1]
		} else if len(key) == 1 && key[0] < 128 {
			m.filterText += key
		} else if key == "space" {
			m.filterText += " "
		}

		m.updateFilterList()
		if m.selectCursor >= len(m.filterList) &&
			len(m.filterList) > 0 {
			m.selectCursor = len(m.filterList) - 1
		}

		return nil
	}

	switch key {
	case "up", "k":
		if m.selectCursor > 0 {
			m.selectCursor--
		}

	case "down", "j":
		if m.selectCursor < len(m.filterList)-1 {
			m.selectCursor++
		}

	case "a":
		if !(len(m.filterList) > 0) {
			return nil
		}

		toggleOn := false
		if !m.filterList[0].selected {
			toggleOn = true
		}

		if toggleOn {
			for _, g := range m.filterList {
				g.selected = true
			}
		} else {
			for _, g := range m.filterList {
				g.selected = false
			}
		}

	case "space":
		group := m.filterList[m.selectCursor]
		group.selected = !group.selected

	case "enter":
		m.setScreen(confirmView)

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
		m.setScreen(selectView)
	}

	return nil
}
