package main

import (
	tea "charm.land/bubbletea/v2"
)

type model struct {
	screen    string
	altscreen bool

	cursor   int
	options  []string
	selected map[int]struct{}
}

func (m *model) setScreen(name string) {
	m.screen = name
	m.altscreen = !(m.screen == "")
}

func initialModel(groups []group) model {
	m := model{
		cursor:   0,
		options:  make([]string, 0, len(groups)),
		selected: make(map[int]struct{}),
	}

	for _, group := range groups {
		m.options = append(m.options, group.name)
	}

	m.setScreen("choose")

	return m
}

func (m model) Init() tea.Cmd {
	return nil
}
