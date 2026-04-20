package main

import (
	tea "charm.land/bubbletea/v2"
)

type model struct {
	screen    string
	altscreen bool

	cursor    int
	selected  map[int]struct{}
	confirmed bool
}

func (m *model) setScreen(name string) {
	m.screen = name
	m.altscreen = !(m.screen == "")
}

func initialModel() model {
	m := model{
		cursor:   0,
		selected: make(map[int]struct{}),
	}

	m.setScreen("choose")

	return m
}

func (m model) Init() tea.Cmd {
	return nil
}
