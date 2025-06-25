package main

import (
	"chai-lang/app/scanner"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Println("Invalid usage\nRun ./run.sh tokenize filename.chai")
		return
	}
	command := args[1]

	if command != "tokenize" {
		fmt.Println("Invalid usage\nRun ./run.sh tokenize filename.chai")
		return
	}

	filename := args[2]
	rawContent, err := os.ReadFile(filename) // command, ./run.sh tokenize test.chai
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	scanner := scanner.NewScanner(rawContent)
	err = scanner.Scan()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Scan completed successfully.")
	}
}
