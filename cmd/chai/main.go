package main

import (
	"chai-lang/internal/ast"
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
	_ = scanner.Scan()
	for _, t := range scanner.GetTokens() {
		fmt.Println(t)
	}
	// Ast printer
	printer := ast.NewAstPrinter()
	expr := ast.BinaryExpr{
		Left: &ast.UnaryExpr{
			Operator: ast.Token{TokenType: ast.MINUS, Lexeme: "-", Literal: nil, Line: 1},
			Right:    &ast.Literal{Value: 123},
		},
		Operator: ast.Token{TokenType: ast.STAR, Lexeme: "*", Literal: nil, Line: 1},
		Right: &ast.Grouping{
			Expression: &ast.Literal{Value: 45.67},
		},
	}
	fmt.Println(printer.Print(&expr))
	exitCode := scanner.GetExitCode()
	os.Exit(exitCode)
}
