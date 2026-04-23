package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	tea "charm.land/bubbletea/v2"
)

var (
	dryRun = flag.Bool("dry-run", false, "Perform a trial run without making actual changes")
	groups []group
)

func main() {
	cdCwd()

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
		for _, g := range groups {
			if g.selected {
				g.install()
			}
		}
	default:
		fmt.Println("Error: Bubbletea model type not model")
	}
}

func cdCwd() {
	wd, err := os.Getwd()
	check(err)
	// For debugger if it creates binary inside ./src
	err = os.Chdir(strings.TrimSuffix(wd, "/src"))
	check(err)
}
