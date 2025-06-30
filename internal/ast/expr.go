package ast

type Expr interface {
	Accept(visitor ExprVisitor) any
}

type ExprVisitor interface {
	VisitBinaryExpr(expr *BinaryExpr) any
	VisitGrouping(expr *Grouping) any
	VisitLiteral(expr *Literal) any
	VisitUnaryExpr(expr *UnaryExpr) any
	VisitBolStmt(stmt *BolStmt) any
	VisitExpressionStmt(stmt *ExpressionStmt) any
	VisitDekhStmt(stmt *DekhStmt) any
	VisitDekh(expr *Dekh) any
}

type BinaryExpr struct {
	Left     Expr
	Operator Token
	Right    Expr
}

func (b *BinaryExpr) Accept(visitor ExprVisitor) any {
	return visitor.VisitBinaryExpr(b)
}

type Grouping struct {
	Expression Expr
}

func (g *Grouping) Accept(visitor ExprVisitor) any {
	return visitor.VisitGrouping(g)
}

type Literal struct {
	Value any
}

func (l *Literal) Accept(visitor ExprVisitor) any {
	return visitor.VisitLiteral(l)
}

type UnaryExpr struct {
	Operator Token
	Right    Expr
}

func (u *UnaryExpr) Accept(visitor ExprVisitor) any {
	return visitor.VisitUnaryExpr(u)
}

type Dekh struct {
	Name Token
}

func (d *Dekh) Accept(visitor ExprVisitor) any {
	return visitor.VisitDekh(d)
}
