package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: gogent <command>")
		return
	}

	command := os.Args[1]

	switch command {
	case "scan":
		scan()
		return
	case "read":
		if len(os.Args) < 3 {
			fmt.Println("file path not provided")
			return
		}
		content, err := readFile(os.Args[2])

		if err != nil {
			fmt.Println("error:", err)
		} else {
			fmt.Println(content)
			return
		}
	default:
		fmt.Println("unknown command", command)
	}
}
