package main

import (
	"flag"

	tea "charm.land/bubbletea/v2"
)

var dryRun = flag.Bool("dry-run", false, "Perform a trial run without making actual changes")

func main() {
	readMappings("src/mappings.conf")

	p := tea.NewProgram(initialModel())
	_, err := p.Run()
	check(err)

	// flag.Parse()
	//
	// fmt.Println()
	// for g := range groups {
	// 	fmt.Println("Info: Symlinking group " + g + ".")
	// 	for _, m := range groups[g] {
	// 		m.createSymlink()
	// 	}
	// }
}
