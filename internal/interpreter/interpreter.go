package interpreter

import (
	"fmt"

	"chai-lang/internal/ast"
	"chai-lang/internal/environment"
	"chai-lang/internal/errors"
)

type Interpreter struct {
	environment *environment.Environment
}

func NewInterpreter() *Interpreter {
	return &Interpreter{
		environment: environment.NewEnvironment(),
	}
}

func (i *Interpreter) VisitLiteral(expr *ast.Literal) any {
	return expr.Value
}

func (i *Interpreter) VisitGrouping(expr *ast.Grouping) any {
	return i.evaluate(expr.Expression)
}

func (i *Interpreter) VisitUnaryExpr(expr *ast.UnaryExpr) any {
	right := i.evaluate(expr.Right)
	switch expr.Operator.TokenType {
	case ast.MINUS:
		i.checkNumberOperand(expr.Operator, right)
		return -right.(float64)
	case ast.ULTA:
		return !i.isTruthy(right)
	default:
		// Unreaacble
		return nil
	}
}

func (i *Interpreter) VisitBinaryExpr(expr *ast.BinaryExpr) any {
	left := i.evaluate(expr.Left)
	right := i.evaluate(expr.Right)

	switch expr.Operator.TokenType {
	case ast.PLUS:
		_, lOk := left.(string)
		_, rOk := right.(string)
		if lOk || rOk {
			return fmt.Sprintf("%v%v", left, right)
		}
		lNumber, lNumOk := left.(float64)
		rNumber, rNumOk := right.(float64)
		if lNumOk && rNumOk {
			return lNumber + rNumber
		}
		panic(errors.NewRuntimeError(expr.Operator, "Operands must be two numbers or two strings for '+' operator"))
	case ast.MINUS:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) - right.(float64)
	case ast.STAR:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) * right.(float64)
	case ast.SLASH:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) / right.(float64)
	case ast.GREATER:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) > right.(float64)
	case ast.GREATER_EQUAL:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) >= right.(float64)
	case ast.LESS:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) < right.(float64)
	case ast.LESS_EQUAL:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) <= right.(float64)
	case ast.EQUAL_EQUAL:
		return i.isEqual(left, right)
	case ast.BANG_EQUAL:
		return !i.isEqual(left, right)
	default:
		// Unreachable code
		return nil
	}
}

func (i *Interpreter) evaluate(expr ast.Expr) any {
	return expr.Accept(i)
}

func (i *Interpreter) VisitExpressionStmt(stmt *ast.ExpressionStmt) any {
	i.evaluate(stmt.Expression)
	return nil
}

func (i *Interpreter) VisitBolStmt(stmt *ast.BolStmt) any {
	value := i.evaluate(stmt.Expression)
	if value == nil {
		value = "khali"
	}
	if value == true {
		value = "haan"
	}
	if value == false {
		value = "nahi"
	}
	fmt.Println(value)
	return nil
}

func (i *Interpreter) VisitDekhStmt(stmt *ast.DekhStmt) any {
	if stmt.Initializer == nil {
		panic(errors.NewRuntimeError(stmt.Name, "Variable declaration must have an initializer"))
	}
	value := i.evaluate(stmt.Initializer)
	i.environment.Define(stmt.Name, value)
	return nil
}

func (i *Interpreter) VisitDekhExpr(expr *ast.DekhExpr) any {
	return i.environment.Get(expr.Name)
}

func (i *Interpreter) VisitAssignExpr(expr *ast.AssignExpr) any {
	value := i.evaluate(expr.Value)
	i.environment.Set(expr.Name, value)
	return nil
}

func (i *Interpreter) VisitBlockStmt(stmt *ast.BlockStmt) any {
	i.executeBlock(stmt.Statements, environment.NewEnclosedEnvironment(i.environment))
	return nil
}

func (i *Interpreter) executeBlock(statements []ast.Stmt, env *environment.Environment) {
	i.environment = env
	for _, stmt := range statements {
		i.execute(stmt)
	}
	i.environment = i.environment.Enclosing()
}

func (i *Interpreter) isTruthy(value any) bool {
	if value == nil {
		return false
	}
	if boolean, ok := value.(bool); ok {
		return boolean
	}
	return true
}

func (i *Interpreter) VisitAgarStmt(stmt *ast.AgarStmt) any {
	condition := i.evaluate(stmt.Condition)
	if i.isTruthy(condition) {
		i.execute(stmt.AgarBranch)
	} else if stmt.WarnaBranch != nil {
		i.execute(stmt.WarnaBranch)
	}
	return nil
}

func (i *Interpreter) VisitJabTakStmt(stmt *ast.JabTakStmt) any {
	for i.isTruthy(i.evaluate(stmt.Condition)) {
		i.execute(stmt.Body)
	}
	return nil
}

func (i *Interpreter) isEqual(left, right any) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if leftType, ok := left.(string); ok {
		if rightType, ok := right.(string); ok {
			return leftType == rightType
		}
		return false
	}
	return left == right
}

func (i *Interpreter) Interpret(statements []ast.Stmt) {
	defer func() {
		if r := recover(); r != nil {
			if runtimeError, ok := r.(*errors.RuntimeError); ok {
				errors.GlobalErrorState.HasRuntimeError = true
				errors.ReportRuntimeError(runtimeError)
			}
		}
	}()
	for _, stmt := range statements {
		i.execute(stmt)
	}
}

func (i *Interpreter) execute(stmt ast.Stmt) {
	stmt.Accept(i)
}

func (i *Interpreter) checkNumberOperand(operator ast.Token, operand any) {
	if _, ok := operand.(float64); !ok {
		panic(errors.NewRuntimeError(operator, "Operand must be a number"))
	}
}

func (i *Interpreter) checkNumberOperands(operator ast.Token, left, right any) {
	if _, ok := left.(float64); !ok {
		panic(errors.NewRuntimeError(
			operator,
			"Left operand must be a number",
		))
	}
	if _, ok := right.(float64); !ok {
		panic(errors.NewRuntimeError(operator, "Right operand must be a number"))
	}
}
