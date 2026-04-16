package main

import (
	"fmt"
)

func main() {
	initDirs()
	readMappings()

	fmt.Println()
	for g := range groups {
		fmt.Println("Info: Symlinking group " + g + ".")
		for _, m := range groups[g] {
			m.createSymlink()
		}
	}
}
