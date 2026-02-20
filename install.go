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

	for _, e := range groups {
		fmt.Println("Info: Symlinking group " + e.name + ".")
		for _, m := range e.mappings {
			m.createSymlink()
		}
	}
}

const MAPPINGS_FILE string = "mappings.conf"

var (
	home_dir    string
	content_dir string
	groups      []*Group
)

func initDirs() {
	home_dir = os.Getenv("HOME")

	wd, err := os.Getwd()
	check(err)
	content_dir = wd + "/content"
}

func readMappings() {
	file_bytes, err := os.ReadFile(MAPPINGS_FILE)
	check(err)
	file_slice := strings.Split(string(file_bytes), "\n")
	file_slice = slices.DeleteFunc(file_slice, func(line string) bool {
		if line == "" {
			return true
		}
		return false
	})

	get_line_type := func(line string) string {
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
	for _, line := range file_slice {
		var line_type string
		line_type = get_line_type(line)

		if line_type == "group" {
			line = strings.Trim(line, "#")
			group = &Group{name: strings.TrimSpace(line)}
			groups = append(groups, group)
		}

		if line_type == "mapping" {
			if group.name == "" {
				fmt.Println("Error: Mapping without a group.")
				os.Exit(1)
			}

			var line_split [2]string = [2]string(strings.Split(line, ":"))

			line_split[0] = strings.TrimSpace(line_split[0])
			line_split[1] = strings.TrimSpace(line_split[1])

			mapping := &Mapping{
				src:  content_dir + "/" + line_split[0],
				dest: strings.ReplaceAll(line_split[1], "~", home_dir),
			}

			group.mappings = append(group.mappings, mapping)
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
	file_info, err := os.Stat(m.src)
	check(err)

	if !(file_info.IsDir()) {
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

	dir_files, err := os.ReadDir(m.src)
	check(err)
	for _, file := range dir_files {
		var new_m Mapping = Mapping{
			src:  m.src + "/" + file.Name(),
			dest: m.dest + "/" + file.Name(),
		}
		new_m.createSymlink()
	}
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
