package main

import (
	"flag"
	"fmt"
)

var dryRun = flag.Bool("dry-run", false, "Perform a trial run without making actual changes")

func main() {
	groups := readMappings("src/mappings.conf")
	flag.Parse()

	fmt.Println()
	for g := range groups {
		fmt.Println("Info: Symlinking group " + g + ".")
		for _, m := range groups[g] {
			m.createSymlink()
		}
	}
}
