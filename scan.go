package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// list of default files/directories that the scan function will ignore
// used hashmap for O(1) look up

func scan() {
	fmt.Println(".")
	var ignored, parserErr = parseGitignore(".gitignore")

	if parserErr != nil {
		return
	}

	Directory := scanDirectory(".", 2, ignored)

	//prints sub directories and files
	for _, line := range Directory {
		fmt.Println(line)
	}
}

// A DFS function to print all the directories and files present recursively
func scanDirectory(path string, indent int, ignored []string) []string {
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
		flag := false
		for _, ignore := range ignored {
			matched, _ := filepath.Match(ignore, entry.Name())
			if matched {
				flag = true
				break
			}
		}
		if flag {
			continue
		}

		indentSpace := strings.Repeat(" ", indent)
		JoinedName := indentSpace + entry.Name()

		lines = append(lines, JoinedName)

		//recursively calls helper again to print the children elements
		if entry.IsDir() {
			joinedPath := filepath.Join(path, entry.Name())
			childLines := scanDirectory(joinedPath, indent+2, ignored)
			lines = append(lines, childLines...)
		}
	}

	//we finally return that lines list
	return lines
}
