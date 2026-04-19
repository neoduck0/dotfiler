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
	for _, g := range groups {
		g.install()
	}
}
