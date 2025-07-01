package errors

import (
	"fmt"

	"chai-lang/internal/ast"
)

type ErrorState struct {
	HasRuntimeError bool
	HasParserError  bool
}

var GlobalErrorState = &ErrorState{}

type RuntimeError struct {
	Message string
	Token   ast.Token
}

func (re *RuntimeError) Error() string {
	return re.Message
}

func NewRuntimeError(token ast.Token, message string) *RuntimeError {
	return &RuntimeError{
		Message: message,
		Token:   token,
	}
}

func ReportRuntimeError(error *RuntimeError) {
	fmt.Printf("%s\n[line %d]\n", error.Error(), error.Token.Line)
}

func ReportError(line int, where string, message string) {
	fmt.Printf("[line %d] Error%s: %s\n", line, where, message)
}

type ParserError struct {
	msg string
}

func NewParserError(msg string) ParserError {
	return ParserError{msg: msg}
}

func (p ParserError) Error() string {
	return p.msg
}

type BreakSignal struct{}

func (b BreakSignal) Error() string {
	return "Loop break signal"
}
