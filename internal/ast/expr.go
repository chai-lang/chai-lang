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
	VisitDekhExpr(expr *DekhExpr) any
	VisitAssignExpr(expr *AssignExpr) any
	VisitBlockStmt(stmt *BlockStmt) any
	VisitAgarStmt(stmt *AgarStmt) any
	VisitJabTakStmt(stmt *JabTakStmt) any
	VisitRehneDeStmt(stmt *RehneDeStmt) any
	VisitAglaStmt(stmt *AglaStmt) any
	VisitCallExpr(expr *CallExpr) any
	VisitKaamStmt(stmt *KaamStmt) any
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

type DekhExpr struct {
	Name Token
}

func (d *DekhExpr) Accept(visitor ExprVisitor) any {
	return visitor.VisitDekhExpr(d)
}

type AssignExpr struct {
	Name  Token
	Value Expr
}

func (a *AssignExpr) Accept(visitor ExprVisitor) any {
	return visitor.VisitAssignExpr(a)
}

type CallExpr struct {
	Callee    Expr
	Paren     Token
	Arguments []Expr
}

func (c *CallExpr) Accept(visitor ExprVisitor) any {
	return visitor.VisitCallExpr(c)
}
