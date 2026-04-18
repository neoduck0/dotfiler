package main

import (
	"maps"
	"slices"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	// "choose" or "output"
	screen   string
	cursor   int
	options  []string
	selected map[int]struct{}
}

func initialModel(groups map[string][]*mapping) model {
	m := model{
		screen:   "choose",
		cursor:   0,
		options:  slices.Collect(maps.Keys(groups)),
		selected: make(map[int]struct{}),
	}
	return m
}

func (m model) Init() tea.Cmd {
	return nil
}
