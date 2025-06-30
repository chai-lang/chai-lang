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
		if leftStr, ok := left.(string); ok {
			if rightStr, ok := right.(string); ok {
				return leftStr + rightStr
			}
			return nil
		}
		return left.(float64) + right.(float64)
	case ast.MINUS:
		return left.(float64) - right.(float64)
	case ast.STAR:
		return left.(float64) * right.(float64)
	case ast.SLASH:
		return left.(float64) / right.(float64)
	case ast.GREATER:
		return left.(float64) > right.(float64)
	case ast.GREATER_EQUAL:
		return left.(float64) >= right.(float64)
	case ast.LESS:
		return left.(float64) < right.(float64)
	case ast.LESS_EQUAL:
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
