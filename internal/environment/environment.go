package environment

import (
	"github.com/chai-lang/chai-lang/internal/ast"
	"github.com/chai-lang/chai-lang/internal/errors"
)

type Environment struct {
	values    map[string]any
	enclosing *Environment
}

func NewEnvironment() *Environment {
	return &Environment{
		values:    make(map[string]any),
		enclosing: nil,
	}
}

func NewEnclosedEnvironment(enclosing *Environment) *Environment {
	return &Environment{
		values:    make(map[string]any),
		enclosing: enclosing,
	}
}

func (env *Environment) Set(name ast.Token, value any) {
	if _, exists := env.values[name.Lexeme]; exists {
		env.values[name.Lexeme] = value
		return
	}
	if env.enclosing != nil {
		env.enclosing.Set(name, value)
		return
	}
	panic(errors.NewRuntimeError(name, "Variable '"+name.Lexeme+"' is not defined."))

}

func (env *Environment) Get(name ast.Token) any {
	if value, exists := env.values[name.Lexeme]; exists {
		return value
	}
	if env.enclosing != nil {
		return env.enclosing.Get(name)
	}
	panic("Variable '" + name.Lexeme + "' is not defined.")
}

func (env *Environment) Define(name ast.Token, value any) {
	if _, exists := env.values[name.Lexeme]; exists {
		panic(errors.NewRuntimeError(name, "Variable '"+name.Lexeme+"' is already defined."))
	}
	env.values[name.Lexeme] = value
}

func (env *Environment) Enclosing() *Environment {
	if env.enclosing == nil {
		panic("No enclosing environment.")
	}
	return env.enclosing
}

func (env *Environment) DefineNative(name string, value any) {
	if _, exists := env.values[name]; exists {
		panic(errors.NewRuntimeError(ast.Token{TokenType: ast.IDENTIFIER, Lexeme: name}, "Native function '"+name+"' is already defined."))
	}
	env.values[name] = value
}
