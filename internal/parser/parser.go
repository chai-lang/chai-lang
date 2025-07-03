package parser

import (
	"fmt"
	"slices"

	"github.com/chai-lang/chai-lang/internal/ast"
	"github.com/chai-lang/chai-lang/internal/errors"
)

type Parser struct {
	tokens  []ast.Token
	current int
}

func NewParser(tokens []ast.Token) *Parser {
	return &Parser{
		tokens:  tokens,
		current: 0,
	}
}

func (p *Parser) expression() ast.Expr {
	return p.assignment()
}

func (p *Parser) assignment() ast.Expr {
	expr := p.equality()
	if p.match(ast.HAI) {
		hai := p.previous()
		value := p.assignment()
		if dekhExpr, ok := expr.(*ast.DekhExpr); ok {
			return &ast.AssignExpr{
				Name:  dekhExpr.Name,
				Value: value,
			}
		}
		p.error(hai, "Invalid assignment target. Expected variable name after 'hai'.")
	}

	return expr
}

func (p *Parser) equality() ast.Expr {
	expr := p.comparison()

	for p.match(ast.BANG_EQUAL, ast.EQUAL_EQUAL) {
		operator := p.previous()
		right := p.comparison()
		expr = &ast.BinaryExpr{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}
	return expr
}

func (p *Parser) match(types ...ast.TokenType) bool {
	if slices.ContainsFunc(types, p.check) {
		p.advance()
		return true
	}
	return false
}

func (p *Parser) check(tokenType ast.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().TokenType == tokenType
}

func (p *Parser) isAtEnd() bool {
	return p.current >= len(p.tokens) || p.tokens[p.current].TokenType == ast.EOF
}

func (p *Parser) advance() ast.Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) previous() ast.Token {
	if p.current == 0 {
		return ast.Token{}
	}
	return p.tokens[p.current-1]
}

func (p *Parser) peek() ast.Token {
	if p.isAtEnd() {
		return ast.Token{TokenType: ast.EOF}
	}
	return p.tokens[p.current]
}

func (p *Parser) comparison() ast.Expr {
	expr := p.term()

	for p.match(ast.GREATER, ast.GREATER_EQUAL, ast.LESS, ast.LESS_EQUAL) {
		operator := p.previous()
		right := p.term()
		expr = &ast.BinaryExpr{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}
	return expr
}

func (p *Parser) term() ast.Expr {
	expr := p.factor()
	for p.match(ast.PLUS, ast.MINUS) {
		operator := p.previous()
		right := p.factor()
		expr = &ast.BinaryExpr{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}
	return expr
}

func (p *Parser) factor() ast.Expr {
	expr := p.unary()
	for p.match(ast.STAR, ast.SLASH) {
		operator := p.previous()
		right := p.unary()
		expr = &ast.BinaryExpr{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}
	return expr
}

func (p *Parser) unary() ast.Expr {
	if p.match(ast.ULTA, ast.MINUS) {
		operator := p.previous()
		right := p.unary()
		return &ast.UnaryExpr{
			Operator: operator,
			Right:    right,
		}
	}
	return p.primary()
}

func (p *Parser) primary() ast.Expr {
	if p.match(ast.IDENTIFIER) {
		return &ast.DekhExpr{Name: p.previous()}
	}
	if p.match(ast.KHALI) {
		return &ast.Literal{Value: nil}
	}
	if p.match(ast.HAAN) {
		return &ast.Literal{Value: true}
	}
	if p.match(ast.NAHI) {
		return &ast.Literal{Value: false}
	}
	if p.match(ast.STRING, ast.NUMBER) {
		return &ast.Literal{Value: p.previous().Literal}
	}

	if p.match(ast.LEFT_PAREN) {
		expr := p.expression()
		p.consume(ast.RIGHT_PAREN, "Expected ')' after expression.")
		return &ast.Grouping{Expression: expr}
	}
	err := errors.NewParserError("Expected expression, found " + p.peek().Lexeme)
	panic(err)
}

func (p *Parser) consume(tokenType ast.TokenType, message string) ast.Token {
	if p.check(tokenType) {
		return p.advance()
	}
	p.error(p.peek(), message)
	// Unreachble
	return ast.Token{}
}

func (p *Parser) error(token ast.Token, message string) {
	var where string
	if token.TokenType == ast.EOF {
		where = " at end"
	} else {
		where = " at '" + token.Lexeme + "'"
	}
	err := errors.NewParserError(fmt.Sprintf("[line %d] Error%s: %s\n", token.Line+1, where, message))
	panic(err)
}

func (p *Parser) parse() []ast.Stmt {
	statements := make([]ast.Stmt, 0)
	for !p.isAtEnd() {
		statements = append(statements, p.declaration())
	}

	return statements
}

func (p *Parser) declaration() ast.Stmt {
	if p.match(ast.DEKH) {
		return p.dekhStatement()
	}
	return p.statement()
}

func (p *Parser) statement() ast.Stmt {
	if p.match(ast.AGAR) {
		return p.agarStatement()
	}
	if p.match(ast.BOL) {
		return p.bolStatement()
	}
	if p.match(ast.LEFT_BRACE) {
		return p.blockStatement()
	}
	if p.match(ast.JAB_TAK) {
		return p.jabtakStatement()
	}
	if p.match(ast.REHNE_DE) {
		return p.rehneDeStatement()
	}
	if p.match(ast.AGLA) {
		return p.aglaStatement()
	}

	return p.expressionStatement()
}

func (p *Parser) agarStatement() ast.Stmt {
	condition := p.expression()
	agarBranch := p.statement()
	var warnaBranch ast.Stmt = nil
	if p.match(ast.WARNA) {
		warnaBranch = p.statement()
	}

	return &ast.AgarStmt{
		Condition:   condition,
		AgarBranch:  agarBranch,
		WarnaBranch: warnaBranch,
	}
}

func (p *Parser) jabtakStatement() ast.Stmt {
	p.consume(ast.LEFT_PAREN, "Expected '(' after 'jabtak'.")
	condition := p.expression()
	p.consume(ast.RIGHT_PAREN, "Expected ')' after condition.")

	p.consume(ast.TAB_TAK, "Expected 'tabtak' after condition in 'jabtak' statement.")
	body := p.statement()

	return &ast.JabTakStmt{
		Condition: condition,
		Body:      body,
	}
}

func (p *Parser) rehneDeStatement() ast.Stmt {
	rehneDeToken := p.previous()
	p.consume(ast.SEMICOLON, "Expected ';' after 'rehnede'")
	return &ast.RehneDeStmt{
		Token: rehneDeToken,
	}
}

func (p *Parser) aglaStatement() ast.Stmt {
	aglaToken := p.previous()
	p.consume(ast.SEMICOLON, "Expected ';' after 'agla'")
	return &ast.AglaStmt{
		Token: aglaToken,
	}
}

func (p *Parser) blockStatement() ast.Stmt {
	statements := make([]ast.Stmt, 0)
	for !p.check(ast.RIGHT_BRACE) && !p.isAtEnd() {
		statements = append(statements, p.declaration())
	}
	p.consume(ast.RIGHT_BRACE, "Expected '}' to close block statement.")
	return &ast.BlockStmt{
		Statements: statements,
	}
}

func (p *Parser) dekhStatement() ast.Stmt {
	name := p.consume(ast.IDENTIFIER, "Expected variable name after 'dekh'.")
	if !p.check(ast.HAI) {
		p.error(p.peek(), "Expected 'hai' after variable name.")
	}
	p.consume(ast.HAI, "Expected 'hai' after variable name.")
	initialization := p.expression()
	if !p.check(ast.SEMICOLON) {
		p.error(p.peek(), "Expected ';' after variable declaration.")
	}
	p.consume(ast.SEMICOLON, "Expected ';' after variable declaration.")
	return &ast.DekhStmt{
		Name:        name,
		Initializer: initialization,
	}
}

func (p *Parser) bolStatement() ast.Stmt {
	expr := p.expression()
	if !p.check(ast.SEMICOLON) {
		p.error(p.peek(), "Expected ';' after expression.")
	}
	p.consume(ast.SEMICOLON, "Expected ';' after expression.")
	return &ast.BolStmt{Expression: expr}
}

func (p *Parser) expressionStatement() ast.Stmt {
	expr := p.expression()
	if !p.check(ast.SEMICOLON) {
		p.error(p.peek(), "Expected ';' after expression.")
	}
	p.consume(ast.SEMICOLON, "Expected ';' after expression.")
	return &ast.ExpressionStmt{Expression: expr}
}

func (p *Parser) Parse() []ast.Stmt {
	defer func() {
		if r := recover(); r != nil {
			if parserErr, ok := r.(errors.ParserError); ok {
				errors.GlobalErrorState.HasParserError = true
				errors.ReportError(p.peek().Line+1, "", parserErr.Error())
			}
		}
	}()
	return p.parse()
}
