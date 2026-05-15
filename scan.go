package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// list of default files/directories that the scan function will ignore
// used hashmap for O(1) look up

var ignored = map[string]bool{
	".git":         true,
	".gocache":     true,
	"node_modules": true,
	"dist":         true,
	"build":        true,
	"vendor":       true,
}

func scan() {
	fmt.Println(".")
	Directory := scanDirectory(".", 2)

	//prints sub directories and files
	for _, line := range Directory {
		fmt.Println(line)
	}
}

// A DFS function to print all the directories and files present recursively
func scanDirectory(path string, indent int) []string {
	entries, err := os.ReadDir(path)

	//If we encounter an error while reading the file directory we simply print that error
	if err != nil {
		fmt.Println("scan error:", err)
		return []string{}
	}

	//initialize lines list which will contain each file/subdirectory path i.e
	//indentation + it's name
	var lines []string

	for _, entry := range entries {

		//we skip the directories or files present in the ignored map
		if ignored[entry.Name()] {
			continue
		}

		indentSpace := strings.Repeat(" ", indent)
		JoinedName := indentSpace + entry.Name()

		lines = append(lines, JoinedName)

		//recursively calls helper again to print the children elements
		if entry.IsDir() {
			joinedPath := filepath.Join(path, entry.Name())
			childLines := scanDirectory(joinedPath, indent+2)
			lines = append(lines, childLines...)
		}
	}

	//we finally return that lines list
	return lines
}
