package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: mini-claude <command>")
		fmt.Println("This is a CLI Coding Agent like Claude Code built in Go")
		return
	}

	command := os.Args[1]

	if command == "scan" {
		scan()
		return
	}

	fmt.Println("unknown command", command)
}
