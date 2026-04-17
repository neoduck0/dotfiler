package main

import (
	"fmt"
)

func main() {
	readMappings("mappings.conf")

	fmt.Println()
	for g := range groups {
		fmt.Println("Info: Symlinking group " + g + ".")
		for _, m := range groups[g] {
			m.createSymlink()
		}
	}
}
