package scanner

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
	BANG          TokenType = "BANG"       // !
	BANG_EQUAL    TokenType = "BANG_EQUAL" // !=
	EQUAL         TokenType = "EQUAL"
	EQUAL_EQUAL   TokenType = "EQUAL_EQUAL"   // ==
	GREATER       TokenType = "GREATER"       // >
	LESS          TokenType = "LESS"          // <
	GREATER_EQUAL TokenType = "GREATER_EQUAL" // >=
	LESS_EQUAL    TokenType = "LESS_EQUAL"    // <=
	SLASH         TokenType = "SLASH"
	NUMBER        TokenType = "NUMBER"
	IDENTIFIER    TokenType = "IDENTIFIER"

	// Keywords
	AND      TokenType = "AND"
	OR       TokenType = "OR"
	IF       TokenType = "IF"
	ELIF     TokenType = "ELIF"
	ELSE     TokenType = "ELSE"
	WHILE    TokenType = "WHILE"
	FOR      TokenType = "FOR"
	FUNCTION TokenType = "FUNCTION"
	RETURN   TokenType = "RETURN"
	PRINT    TokenType = "PRINT"
	TRUE     TokenType = "TRUE"
	FALSE    TokenType = "FALSE"
	NIL      TokenType = "NIL"
)

type Scanner struct {
	source []byte
}

func NewScanner(source []byte) *Scanner {
	return &Scanner{
		source: source,
	}
}

func (s *Scanner) Scan() error {
	return nil
}
