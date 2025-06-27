package main

import (
	"chai-lang/internal/ast"
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
	// TODO: Handle errors
	parser := parser.NewParser(scanner.GetTokens())
	tree := parser.Parse()
	if tree == nil {
		printer := ast.NewAstPrinter()
		printer.Print(tree)
	}
	exitCode := scanner.GetExitCode()
	os.Exit(exitCode)
}
