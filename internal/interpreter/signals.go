package interpreter

import "github.com/chai-lang/chai-lang/internal/ast"

type BreakSignal struct {
	Token ast.Token
}

func (b BreakSignal) Error() string {
	return "Break signal received"
}

type ContinueSignal struct {
	Token ast.Token
}

func (c ContinueSignal) Error() string {
	return "Continue signal received"
}

type ReturnSignal struct {
	Token ast.Token
	Value any
}

func (r ReturnSignal) Error() string {
	return "Return signal received"
}
