package main

import (
	"flag"
	"fmt"

	tea "charm.land/bubbletea/v2"
)

var (
	dryRun = flag.Bool("dry-run", false, "Perform a trial run without making actual changes")
	groups []group
)

func main() {
	readMappings("src/mappings.conf")
	flag.Parse()

	p := tea.NewProgram(initialModel())
	m, err := p.Run()
	check(err)

	switch m := m.(type) {
	case model:
		if !m.confirmed {
			fmt.Println("Info: Gracefully quiting")
			return
		}
		for i := range m.selected {
			groups[i].install()
		}
	default:
		fmt.Println("Error: Bubbletea model type not model")
	}
}
