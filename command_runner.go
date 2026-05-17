package main

import (
	"fmt"
	"os/exec"
	"strings"
)

var blockedCmds = []string{
	".env",
	"id_rsa",
	"credentials",
}

var dangerousCmds = []string{
	"rm -rf",
	"git reset --hard",
	"git push --force",
	"git push -f",
}

func runCommand(command string) (string, error) {

	//empty command check
	if strings.TrimSpace(command) == "" {
		return "", fmt.Errorf("empty command")
	}

	//blocked command check
	check := isBlocked(command)
	if check {
		return "", fmt.Errorf("the command has been blocked")
	}

	//dangerous command check
	check = isDangerous(command)
	if check {
		fmt.Println("Warning: proceed carefully, it's a dangerous command")
	}

	//show command and ask for input
	fmt.Println(command, ", do you want to proceed [y/n]: ")
	var input string
	fmt.Scanln(&input)

	//execute command if user inputs 'y'
	if input == "y" {
		parts := strings.Fields(command)
		cmd := exec.Command(parts[0], parts[1:]...)
		out, err := cmd.Output()
		if err != nil {
			return "", err
		}
		return string(out), nil
	} else if input == "n" {
		return "", fmt.Errorf("command rejected by user")
	} else {
		return "", fmt.Errorf("invalid input for approval")
	}
}

// helper function to check if a command belongs to blocked list
func isBlocked(command string) bool {
	for _, blockcmd := range blockedCmds {
		if strings.Contains(command, blockcmd) {
			return true
		}
	}
	return false
}

// helper function to check if a command belongs to dangerous list
func isDangerous(command string) bool {
	for _, dangercmd := range dangerousCmds {
		if strings.Contains(command, dangercmd) {
			return true
		}
	}
	return false
}
