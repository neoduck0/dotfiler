package main

import (
	"strings"

	"charm.land/bubbles/v2/textinput"
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

	filterInput textinput.Model
	filterList  []*group

	confirmed bool
}

func (m *model) updateFilterList() {
	m.filterList = m.filterList[:0]

	if m.filterInput.Value() == "" {
		for i := range groups {
			m.filterList = append(m.filterList, &groups[i])
		}
		return
	}

	for i := range groups {
		if strings.Contains(groups[i].name, m.filterInput.Value()) {
			m.filterList = append(m.filterList, &groups[i])
		}
	}
}

func initialModel() *model {
	m := model{
		currentView:  selectView,
		selectCursor: 0,
		filterInput:  textinput.New(),
		filterList:   make([]*group, 0, len(groups)),
	}

	m.filterInput.Prompt = ""

	m.updateFilterList()

	return &m
}

func (m *model) Init() tea.Cmd {
	return nil
}
