package parser

import (
	"chai-lang/internal/ast"
	"chai-lang/internal/errors"
	"fmt"
	"slices"
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
	return p.equality()
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
		statements = append(statements, p.statement())
	}

	return statements
}

func (p *Parser) statement() ast.Stmt {
	if p.match(ast.BOL) {
		return p.bolStatement()
	}
	return p.expressionStatement()
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
