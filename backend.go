package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

type mapping struct {
	src  string
	dest string
}

func (m mapping) createSymlink() {
	fileInfo, err := os.Stat(m.src)
	check(err)

	if !(fileInfo.IsDir()) {
		err = os.MkdirAll(filepath.Dir(m.dest), 0o755)
		if !os.IsExist(err) {
			check(err)
		}
		err = os.Remove(m.dest)
		if !os.IsNotExist(err) {
			check(err)
		}
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

var groups map[string][]*mapping = make(map[string][]*mapping)

func readMappings(mappingsFile string) {
	fileBytes, err := os.ReadFile(mappingsFile)
	check(err)
	fileSlice := strings.Split(string(fileBytes), "\n")
	fileSlice = slices.DeleteFunc(fileSlice, func(line string) bool {
		if line == "" {
			return true
		}
		return false
	})

	getLineType := func(line string) string {
		if strings.Contains(line, "[") {
			return "group"
		} else if strings.Contains(line, ":") {
			return "mapping"
		} else {
			fmt.Println("Error: Bad line in mapping.conf file.")
			os.Exit(1)
			return ""
		}
	}

	var currentGroup string
	var homeDir string
	var contentDir string
	initDirs := func() {
		homeDir = os.Getenv("HOME")

		wd, err := os.Getwd()
		check(err)
		contentDir = wd + "/content"
	}
	initDirs()

	for _, line := range fileSlice {
		var lineType string
		lineType = getLineType(line)

		if lineType == "group" {
			line = strings.Trim(line, "[]")
			currentGroup = strings.TrimSpace(line)
		}

		if lineType == "mapping" {
			if currentGroup == "" {
				fmt.Println("Error: Mapping \"" + line + "\" is without a group.")
				os.Exit(1)
			}

			var lineArr [2]string = [2]string(strings.Split(line, ":"))

			lineArr[0] = strings.TrimSpace(lineArr[0])
			lineArr[1] = strings.TrimSpace(lineArr[1])

			mapping := &mapping{
				src:  contentDir + "/" + lineArr[0],
				dest: strings.ReplaceAll(lineArr[1], "~", homeDir),
			}

			groups[currentGroup] = append(groups[currentGroup], mapping)
		}
	}
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
