package main

import (
	"os/exec"
)

func search(pattern string, path string) (string, error) {
	args := []string{"-rn", pattern, path}

	//we need to get the ignoreList so that search excludes those directories
	ignoreList, parserErr := parseGitignore(".gitignore")
	if parserErr != nil {
		return "", parserErr
	}

	//loop through each item in the ignored list and appedn it to args with exclude flag
	for _, item := range ignoreList {
		flag := "--exclude-dir=" + item
		args = append(args, flag)
	}

	//execute the grep commands with the arguments we have built
	cmd := exec.Command("grep", args...)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}
