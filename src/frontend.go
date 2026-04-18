package main

import (
	"maps"
	"slices"

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

func initialModel(groups map[string][]*mapping) model {
	m := model{
		cursor:   0,
		options:  slices.Collect(maps.Keys(groups)),
		selected: make(map[int]struct{}),
	}
	m.setScreen("choose")
	return m
}

func (m model) Init() tea.Cmd {
	return nil
}
