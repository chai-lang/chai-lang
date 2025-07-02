package interpreter

import "chai-lang/internal/ast"

type BreakSignal struct {
	Token ast.Token
}

func (b BreakSignal) Error() string {
	return "Break signal received"
}

type ContinueSignal struct{}

func (c ContinueSignal) Error() string {
	return "Continue signal received"
}
