package ast

type Stmt interface {
	Accept(visitor ExprVisitor) any
}

type BolStmt struct {
	Expression Expr
}

func (b *BolStmt) Accept(visitor ExprVisitor) any {
	return visitor.VisitBolStmt(b)
}

type ExpressionStmt struct {
	Expression Expr
}

func (e *ExpressionStmt) Accept(visitor ExprVisitor) any {
	return visitor.VisitExpressionStmt(e)
}

type DekhStmt struct { // variable declaration
	Name        Token
	Initializer Expr
}

func (d *DekhStmt) Accept(visitor ExprVisitor) any {
	return visitor.VisitDekhStmt(d)
}
