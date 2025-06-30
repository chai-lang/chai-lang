package interpreter

import "chai-lang/internal/ast"

type Interpreter struct {
}

func NewInterpreter() *Interpreter {
	return &Interpreter{}
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
		lString, lOk := left.(string)
		rString, rOk := right.(string)
		if lOk && rOk {
			return lString + rString
		}
		lNumber, lNumOk := left.(float64)
		rNumber, rNumOk := right.(float64)
		if lNumOk && rNumOk {
			return lNumber + rNumber
		}
		panic(NewRuntimeError(expr.Operator, "Operands must be two numbers or two strings for '+' operator"))
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

func (i *Interpreter) isTruthy(value any) bool {
	if value == nil {
		return false
	}
	if boolean, ok := value.(bool); ok {
		return boolean
	}
	return true
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

func (i *Interpreter) Interpret(tree ast.Expr) any {
	return i.evaluate(tree)
}

func (i *Interpreter) checkNumberOperand(operator ast.Token, operand any) {
	if _, ok := operand.(float64); !ok {
		panic(NewRuntimeError(operator, "Operand must be a number"))
	}
}

func (i *Interpreter) checkNumberOperands(operator ast.Token, left, right any) {
	if _, ok := left.(float64); !ok {
		panic(&RuntimeError{
			Message: "Left operand must be a number",
			Token:   operator,
		})
	}
	if _, ok := right.(float64); !ok {
		panic(NewRuntimeError(operator, "Right operand must be a number"))
	}
}

type RuntimeError struct {
	Message string
	Token   ast.Token
}

func (re *RuntimeError) Error() string {
	return re.Message + " at " + re.Token.Lexeme
}

func NewRuntimeError(token ast.Token, message string) *RuntimeError {
	return &RuntimeError{
		Message: message,
		Token:   token,
	}
}
