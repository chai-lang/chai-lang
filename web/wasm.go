//go:build js && wasm
// +build js,wasm

package main

import (
	"fmt"
	"syscall/js"

	"github.com/chai-lang/chai-lang/internal/errors"
	"github.com/chai-lang/chai-lang/internal/interpreter"
	"github.com/chai-lang/chai-lang/internal/parser"
	"github.com/chai-lang/chai-lang/internal/scanner"
)

func main() {
	registerCallbacks()
	select {}
}

func registerCallbacks() {
	js.Global().Set("run", js.FuncOf(run))
	js.Global().Call("postMessage", "chai-lang:ready", "*")
}

func run(this js.Value, p []js.Value) any {
	if len(p) > 1 {
		return js.Global().Get("Error").New("Too many arguments.")
	}
	if len(p) == 0 {
		return js.Global().Get("Error").New("No code provided.")
	}
	code := p[0].String()
	if len(code) == 0 {
		return js.Global().Get("Error").New("Empty code provided.")
	}
	codeSlice := []byte(code)
	runCodeInteral(codeSlice)
	return js.ValueOf(code)
}

func runCodeInteral(code []byte) js.Value {
	scanner := scanner.NewScanner(code)
	if err := scanner.Scan(); err != nil {
		fmt.Println("Scanner error:", err)
		return js.Null()
	}

	parser := parser.NewParser(scanner.GetTokens())
	tree := parser.Parse()
	if errors.GlobalErrorState.HasParserError {
		fmt.Println("Parser errors detected. Exiting.")
		return js.Null()
	}

	if tree != nil {
		interpreter := interpreter.NewInterpreter()
		interpreter.Interpret(tree)
	}
	return js.Null()
}
