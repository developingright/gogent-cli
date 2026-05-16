package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func readFile(path string) (string, error) {
	file, err := os.ReadFile(path)

	if errors.Is(err, os.ErrNotExist) {
		return "", err
	} else if err != nil {
		return "", err
	}

	if len(file) > 100*1024 {
		return "", fmt.Errorf("file too large")
	}

	lines := strings.Split(string(file), "\n")

	// add line numbers to each line
	for idx, line := range lines {
		lines[idx] = fmt.Sprintf("%d: %s", idx+1, line)
	}

	return strings.Join(lines, "\n"), nil
}
