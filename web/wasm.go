//go:build js && wasm
// +build js,wasm

package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Hello, WebAssembly!")
	registerCallbacks()
	select {}
}

func registerCallbacks() {
	js.Global().Set("run", js.FuncOf(run))
	js.Global().Call("postMessage", "chai-lang:ready", "*")
}

func run(this js.Value, p []js.Value) any {
	arr := []int{1, 2, 3, 4, 5, 6}
	jsArr := js.Global().Get("Array").New()
	for _, v := range arr {
		jsArr.Call("push", v)
	}
	return jsArr
}
