package main

import (
	"flag"
)

var dryRun = flag.Bool("dry-run", false, "Perform a trial run without making actual changes")

func main() {
	groups := readMappings("src/mappings.conf")
	flag.Parse()

	for _, g := range groups {
		g.install()
	}
}
