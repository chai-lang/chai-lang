package interpreter

import (
	"github.com/chai-lang/chai-lang/internal/ast"
	"github.com/chai-lang/chai-lang/internal/environment"
)

type Kaam struct {
	declaration *ast.KaamStmt
}

func NewKaam(declaration *ast.KaamStmt) *Kaam {
	return &Kaam{
		declaration: declaration,
	}
}

func (k *Kaam) Call(interpreter *Interpreter, args []any) (returnValue any) {
	env := environment.NewEnclosedEnvironment(interpreter.environment)
	for i, param := range k.declaration.Parameters {
		env.Define(param, args[i])
	}

	defer func() {
		if r := recover(); r != nil {
			if returnSignal, ok := r.(ReturnSignal); ok {
				returnValue = returnSignal.Value
				return
			}
			panic(r)
		}
	}()

	interpreter.executeBlock(k.declaration.Body.Statements, env)
	return
}

func (k *Kaam) Arity() int {
	return len(k.declaration.Parameters)
}

func (k *Kaam) String() string {
	return "<kaam " + k.declaration.Name.Lexeme + ">"
}
