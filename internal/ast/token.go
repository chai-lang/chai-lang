package ast

import "fmt"

type TokenType string

const (
	// Single character
	LEFT_PAREN  TokenType = "LEFT_PAREN"
	RIGHT_PAREN TokenType = "RIGHT_PAREN"
	LEFT_BRACE  TokenType = "LEFT_BRACE"
	RIGHT_BRACE TokenType = "RIGHT_BRACE"
	COMMA       TokenType = "COMMA"
	SEMICOLON   TokenType = "SEMICOLON"
	DOT         TokenType = "DOT"
	PLUS        TokenType = "PLUS"
	MINUS       TokenType = "MINUS"
	STAR        TokenType = "STAR"
	EOF         TokenType = "EOF"

	// One or two
	BANG_EQUAL    TokenType = "BANG_EQUAL"    // !=
	EQUAL_EQUAL   TokenType = "EQUAL_EQUAL"   // ==
	GREATER       TokenType = "GREATER"       // >
	LESS          TokenType = "LESS"          // <
	GREATER_EQUAL TokenType = "GREATER_EQUAL" // >=
	LESS_EQUAL    TokenType = "LESS_EQUAL"    // <=
	SLASH         TokenType = "SLASH"
	IDENTIFIER    TokenType = "IDENTIFIER"

	// Keywords
	CHAIBANA   TokenType = "CHAIBANA"   // entrypoint
	CHAIKHATAM TokenType = "CHAIKHATAM" // end
	DEKH       TokenType = "DEKH"       // declare
	HAI        TokenType = "HAI"        // assignment
	BOL        TokenType = "BOL"        // print
	SUN        TokenType = "SUN"        // input
	AGAR       TokenType = "AGAR"       // if
	NAHI_TO    TokenType = "NAHI_TOH"   // else if
	WARNA      TokenType = "WARNA"      // else
	JAB_TAK    TokenType = "JAB_TAK"    // while
	TAB_TAK    TokenType = "TAB_TAK"    // while-then
	HAR        TokenType = "HAR"        // for
	ME         TokenType = "ME"         // in
	KAAM       TokenType = "KAAM"       // function
	KAR        TokenType = "KAR"        // call function
	YE_LE      TokenType = "YE_LE"      // return
	HAAN       TokenType = "HAAN"       // true
	NAHI       TokenType = "NAHI"       // false
	REHNE_DE   TokenType = "REHNE_DE"   // BREAK
	KHALI      TokenType = "KHALI"      // nil
	ULTA       TokenType = "ULTA"       // bang (!)

	// Scanned Types
	STRING TokenType = "STRING"
	NUMBER TokenType = "NUMBER"
)

type Token struct {
	TokenType TokenType
	Lexeme    string
	Literal   any
	Line      int
}

func NewToken(tokenType TokenType, lexeme string, literal any, line int) Token {
	return Token{
		TokenType: tokenType,
		Lexeme:    lexeme,
		Literal:   literal,
		Line:      line,
	}
}

func (t Token) String() string {
	return fmt.Sprintf("%s %s %v ", t.TokenType, t.Lexeme, t.Literal)
}
