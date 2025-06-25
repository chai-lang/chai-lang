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
	rawContent, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	scanner := scanner.NewScanner(rawContent)
	_ = scanner.Scan()
	for _, t := range scanner.GetTokens() {
		fmt.Println(t)
	}

	exitCode := scanner.GetExitCode()
	os.Exit(exitCode)
}
