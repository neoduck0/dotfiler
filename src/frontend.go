package main

import (
	"strings"

	tea "charm.land/bubbletea/v2"
)

type view int

const (
	selectView view = iota
	confirmView
)

type model struct {
	currentView view

	selectCursor int

	filterText string
	filterMode bool
	filterList []*group

	confirmed bool
}

func (m *model) updateFilterList() {
	m.filterList = m.filterList[:0]

	if m.filterText == "" {
		for i := range groups {
			m.filterList = append(m.filterList, &groups[i])
		}
		return
	}

	for i := range groups {
		if strings.Contains(groups[i].name, m.filterText) {
			m.filterList = append(m.filterList, &groups[i])
		}
	}
}

func initialModel() *model {
	m := model{
		currentView:  selectView,
		selectCursor: 0,
		filterText:   "",
		filterMode:   false,
		filterList:   make([]*group, 0, len(groups)),
	}

	m.updateFilterList()

	return &m
}

func (m *model) Init() tea.Cmd {
	return nil
}
