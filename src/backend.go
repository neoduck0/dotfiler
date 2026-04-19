package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

type group struct {
	name     string
	mappings []mapping
}

type mapping struct {
	src  string
	dest string
}

func (g group) install() {
	fmt.Println("Info: Symlinking " + g.name + " group")
	for _, m := range g.mappings {
		m.createSymlink()
	}
}

func (m mapping) createSymlink() {
	fileInfo, err := os.Stat(m.src)
	check(err)

	if !(fileInfo.IsDir()) {
		fmt.Println("\tInfo: Found " + m.src)

		if *dryRun {
			return
		}

		fmt.Println("\tInfo: Making path directories for " +
			filepath.Dir(m.dest))
		err = os.MkdirAll(filepath.Dir(m.dest), 0o755)
		if !os.IsExist(err) {
			check(err)
		}

		fmt.Println("\tInfo: Removing  " + m.dest)
		err = os.Remove(m.dest)
		if !os.IsNotExist(err) {
			check(err)
		}

		fmt.Println("\tInfo: Symlinking to" + m.dest)
		err = os.Symlink(m.src, m.dest)
		check(err)
		return
	}

	dirFile, err := os.ReadDir(m.src)
	check(err)
	for _, file := range dirFile {
		var newMapping mapping = mapping{
			src:  m.src + "/" + file.Name(),
			dest: m.dest + "/" + file.Name(),
		}
		newMapping.createSymlink()
	}
}

func readMappings(mappingsFile string) []group {
	fileBytes, err := os.ReadFile(mappingsFile)
	check(err)
	fileSlice := strings.Split(string(fileBytes), "\n")
	fileSlice = slices.DeleteFunc(fileSlice, func(line string) bool {
		if line == "" {
			return true
		}
		return false
	})

	homeDir := os.Getenv("HOME")
	wd, err := os.Getwd()
	check(err)
	contentDir := wd + "/content"

	getLineType := func(line string) string {
		if strings.Contains(line, "[") {
			return "group"
		} else if strings.Contains(line, ":") {
			return "mapping"
		} else {
			fmt.Println("Error: Bad line in mapping.conf file")
			os.Exit(1)
			return ""
		}
	}

	groups := make([]group, 0, 20)
	currentGroup := group{name: ""}
	for _, line := range fileSlice {
		var lineType string
		lineType = getLineType(line)

		if lineType == "group" {
			if currentGroup.name != "" {
				groups = append(groups, currentGroup)
			}

			line = strings.Trim(line, "[]")
			currentGroup = group{name: strings.TrimSpace(line)}
		}

		if lineType == "mapping" {
			if currentGroup.name == "" {
				fmt.Println("Error: Mapping \"" + line + "\" is without a group")
				os.Exit(1)
			}

			var lineArr [2]string = [2]string(strings.Split(line, ":"))

			lineArr[0] = strings.TrimSpace(lineArr[0])
			lineArr[1] = strings.TrimSpace(lineArr[1])

			mapping := mapping{
				src:  contentDir + "/" + lineArr[0],
				dest: strings.ReplaceAll(lineArr[1], "~", homeDir),
			}

			currentGroup.mappings = append(currentGroup.mappings, mapping)
		}
	}
	return groups
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
