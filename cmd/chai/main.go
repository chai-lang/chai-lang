package main

import (
	"chai-lang/internal/errors"
	"chai-lang/internal/interpreter"
	"chai-lang/internal/parser"
	"chai-lang/internal/scanner"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Println("Invalid usage\nUse tokenize command")
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
	if err := scanner.Scan(); err != nil {
		fmt.Fprintf(os.Stderr, "Error scanning file: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Scanned Tokens:")
	for _, token := range scanner.GetTokens() {
		fmt.Println(token)
	}

	parser := parser.NewParser(scanner.GetTokens())
	tree := parser.Parse()
	if errors.GlobalErrorState.HasParserError {
		os.Exit(1)
	}

	if tree != nil {
		fmt.Println("Interpreting the AST:")
		interpreter := interpreter.NewInterpreter()
		interpreter.Interpret(tree)
	}
	exitCode := scanner.GetExitCode()
	os.Exit(exitCode)
}
