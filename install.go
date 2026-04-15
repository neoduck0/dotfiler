package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func main() {
	initDirs()
	readMappings()

	fmt.Println()
	for _, e := range groups {
		fmt.Println("Info: Symlinking group " + e.name + ".")
		for _, m := range e.mappings {
			m.createSymlink()
		}
	}
}

type Group struct {
	name     string
	mappings []*Mapping
}

type Mapping struct {
	src  string
	dest string
}

func (m Mapping) createSymlink() {
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
		var newMapping Mapping = Mapping{
			src:  m.src + "/" + file.Name(),
			dest: m.dest + "/" + file.Name(),
		}
		newMapping.createSymlink()
	}
}

const mappingsFile string = "mappings.conf"

var (
	homeDir    string
	contentDir string
	groups     []*Group
)

func initDirs() {
	homeDir = os.Getenv("HOME")

	wd, err := os.Getwd()
	check(err)
	contentDir = wd + "/content"
}

func readMappings() {
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
		if strings.Contains(line, "#") {
			return "group"
		} else if strings.Contains(line, ":") {
			return "mapping"
		} else {
			fmt.Println("Error: Bad line in mapping.conf file.")
			os.Exit(1)
			return ""
		}
	}

	var group *Group
	for _, line := range fileSlice {
		var lineType string
		lineType = getLineType(line)

		if lineType == "group" {
			line = strings.Trim(line, "#")
			group = &Group{name: strings.TrimSpace(line)}
			groups = append(groups, group)
		}

		if lineType == "mapping" {
			if group.name == "" {
				fmt.Println("Error: Mapping without a group.")
				os.Exit(1)
			}

			var lineArr [2]string = [2]string(strings.Split(line, ":"))

			lineArr[0] = strings.TrimSpace(lineArr[0])
			lineArr[1] = strings.TrimSpace(lineArr[1])

			mapping := &Mapping{
				src:  contentDir + "/" + lineArr[0],
				dest: strings.ReplaceAll(lineArr[1], "~", homeDir),
			}

			group.mappings = append(group.mappings, mapping)
		}
	}
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
