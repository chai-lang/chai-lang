package interpreter

type ChaiCallable interface {
	Call(interpreter *Interpreter, args []any) any
	Arity() int
}
