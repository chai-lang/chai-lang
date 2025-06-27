package ast

import (
	"fmt"
	"strings"
)

type AstPrinter struct{}

func NewAstPrinter() *AstPrinter {
	return &AstPrinter{}
}

func (p *AstPrinter) Print(expr Expr) string {
	return expr.Accept(p).(string)
}

func (p *AstPrinter) VisitBinaryExpr(expr *BinaryExpr) any {
	return p.parenthesize(expr.Operator.Lexeme, expr.Left, expr.Right)
}

func (p *AstPrinter) VisitGrouping(expr *Grouping) any {
	return p.parenthesize("group", expr.Expression)
}

func (p *AstPrinter) VisitLiteral(expr *Literal) any {
	if expr.Value == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", expr.Value)
}

func (p *AstPrinter) VisitUnaryExpr(expr *UnaryExpr) any {
	return p.parenthesize(expr.Operator.Lexeme, expr.Right)
}

func (p *AstPrinter) parenthesize(name string, exprs ...Expr) string {
	var builder strings.Builder
	builder.WriteString("(")
	builder.WriteString(name)
	for _, expr := range exprs {
		builder.WriteString(" ")
		builder.WriteString(expr.Accept(p).(string))
	}
	builder.WriteString(")")
	return builder.String()
}
