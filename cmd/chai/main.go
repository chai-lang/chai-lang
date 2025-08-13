package main

import (
	"fmt"
	"os"

	"github.com/chai-lang/chai-lang/internal/errors"
	"github.com/chai-lang/chai-lang/internal/interpreter"
	"github.com/chai-lang/chai-lang/internal/parser"
	"github.com/chai-lang/chai-lang/internal/scanner"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Invalid usage: chai <filename>")
		return
	}

	filename := args[1]
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

	parser := parser.NewParser(scanner.GetTokens())
	tree := parser.Parse()
	if errors.GlobalErrorState.HasParserError {
		os.Exit(1)
	}

	if tree != nil {
		interpreter := interpreter.NewInterpreter()
		interpreter.Interpret(tree)
	}
	exitCode := scanner.GetExitCode()
	os.Exit(exitCode)
}
