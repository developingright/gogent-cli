package main

import (
	"errors"
	"os"
	"strings"
)

var hardcodedIgnore = []string{
	".gocache",
	".git",
	"node_modules",
	"build",
	"dist",
	"vendor",
}

// function to parse the gitignore file in order to curate ignoreList while scanning repo
// fails gracefully if the file is not found and merges hardcodedIgnore list as well
func parseGitignore(path string) ([]string, error) {

	file, err := os.ReadFile(path)

	if errors.Is(err, os.ErrNotExist) {
		return []string{}, nil
	} else if err != nil {
		return []string{}, err
	}

	//create an array of strings where each string is a line from .gitignore
	lines := strings.Split(string(file), "\n")

	//looping through each line and skipping it if it's blank or a comment
	// otherwise adding it to the list
	var ignoreList []string
	for _, line := range lines {

		line = strings.TrimSuffix(line, "/")

		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}
		ignoreList = append(ignoreList, line)
	}
	ignoreList = append(ignoreList, hardcodedIgnore...)
	return ignoreList, nil
}
