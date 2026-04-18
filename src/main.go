package main

import (
	"flag"

	tea "charm.land/bubbletea/v2"
)

var dryRun = flag.Bool("dry-run", false, "Perform a trial run without making actual changes")

func main() {
	groups := readMappings("src/mappings.conf")
	flag.Parse()

	p := tea.NewProgram(initialModel(groups))
	_, err := p.Run()
	check(err)
}
