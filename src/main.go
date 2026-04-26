package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	tea "charm.land/bubbletea/v2"
)

var (
	dryRun = flag.Bool("dry-run", false, "Perform a trial run without making actual changes")
	groups []group
)

func main() {
	goToRoot()

	readMappings("content/mappings.conf")
	flag.Parse()

	p := tea.NewProgram(initialModel())
	m, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}

	switch m := m.(type) {
	case *model:
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

func goToRoot() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// For debugger if it creates binary inside ./src
	err = os.Chdir(strings.TrimSuffix(wd, "/src"))
	if err != nil {
		log.Fatal(err)
	}
}
