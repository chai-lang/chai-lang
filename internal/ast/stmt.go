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

type BlockStmt struct {
	Statements []Stmt
}

func (b *BlockStmt) Accept(visitor ExprVisitor) any {
	return visitor.VisitBlockStmt(b)
}

type AgarStmt struct {
	Condition   Expr
	AgarBranch  Stmt // if
	WarnaBranch Stmt // else
}

func (a *AgarStmt) Accept(visitor ExprVisitor) any {
	return visitor.VisitAgarStmt(a)
}

type JabTakStmt struct {
	Condition Expr
	Body      Stmt
}

func (j *JabTakStmt) Accept(visitor ExprVisitor) any {
	return visitor.VisitJabTakStmt(j)
}
