package environment

import (
	"chai-lang/internal/ast"
	"chai-lang/internal/errors"
)

type Environment struct {
	values map[string]any
}

func NewEnvironment() *Environment {
	return &Environment{
		values: make(map[string]any),
	}
}

func (env *Environment) Set(name ast.Token, value any) {
	if _, exists := env.values[name.Lexeme]; exists {
		env.values[name.Lexeme] = value
	} else {
		panic(errors.NewRuntimeError(name, "Variable '"+name.Lexeme+"' is not defined."))
	}
}

func (env *Environment) Get(name ast.Token) any {
	if value, exists := env.values[name.Lexeme]; exists {
		return value
	}
	panic("Variable '" + name.Lexeme + "' is not defined.")
}

func (env *Environment) Define(name ast.Token, value any) {
	if _, exists := env.values[name.Lexeme]; exists {
		panic("Variable '" + name.Lexeme + "' is already defined.")
	}
	env.values[name.Lexeme] = value
}
