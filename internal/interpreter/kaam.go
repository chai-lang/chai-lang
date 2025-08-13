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

func (k *Kaam) Call(interpreter *Interpreter, args []any) any {
	env := environment.NewEnclosedEnvironment(interpreter.environment)
	for i, param := range k.declaration.Parameters {
		env.Define(param, args[i])
	}
	interpreter.executeBlock(k.declaration.Body.Statements, env)
	return nil
}

func (k *Kaam) Arity() int {
	return len(k.declaration.Parameters)
}

func (k *Kaam) String() string {
	return "<kaam " + k.declaration.Name.Lexeme + ">"
}
