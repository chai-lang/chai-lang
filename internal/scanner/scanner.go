package scanner

import (
	"chai-lang/internal/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

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

	// Scanned Types
	STRING TokenType = "STRING"
	NUMBER TokenType = "NUMBER"
)

func singleCharacters(c rune) TokenType {
	chars := map[rune]TokenType{
		'{': LEFT_BRACE,
		'}': RIGHT_BRACE,
		'(': LEFT_PAREN,
		')': RIGHT_PAREN,
		',': COMMA,
		';': SEMICOLON,
		'.': DOT,
		'+': PLUS,
		'-': MINUS,
		'*': STAR,
		'/': SLASH,
	}
	return chars[c]
}

func matchOperators(op string) TokenType {
	operators := map[string]TokenType{
		"==": EQUAL_EQUAL,
		"<=": LESS_EQUAL,
		">=": GREATER_EQUAL,
		">":  GREATER,
		"<":  LESS,
		"!=": BANG_EQUAL,
	}
	return operators[op]
}

func matchKeywords(s string) TokenType {
	keywords := map[string]TokenType{
		"hai":        HAI,
		"chaibana":   CHAIBANA,
		"chaikhatam": CHAIKHATAM,
		"dekh":       DEKH,
		"bol":        BOL,
		"sun":        SUN,
		"agar":       AGAR,
		"nahito":     NAHI_TO,
		"warna":      WARNA,
		"jabtak":     JAB_TAK,
		"tabtak":     TAB_TAK,
		"har":        HAR,
		"me":         ME,
		"kaam":       KAAM,
		"kar":        KAR,
		"haan":       HAAN,
		"nahi":       NAHI,
		"khali":      KHALI,
		"rehne_de":   REHNE_DE,
		"yele":       YE_LE,
	}

	return keywords[s]
}

type Scanner struct {
	source   []byte
	tokens   []Token
	start    int
	current  int
	line     int
	exitCode int
}

func NewScanner(source []byte) *Scanner {
	return &Scanner{
		source: source,
		line:   1,
	}
}

func (s *Scanner) Scan() error {
	for !s.isAtEnd() {
		s.scanToken()
	}
	return nil
}

func (s Scanner) GetExitCode() int {
	return s.exitCode
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) advance() (rune, int) {
	if s.isAtEnd() {
		return 0, 0
	}
	r, size := utf8.DecodeRune(s.source[s.current:])
	s.current += size
	return r, size
}

func (s *Scanner) scanToken() {
	s.start = s.current
	r, _ := s.advance()
	switch r {
	case ' ', '\r', '\t':
		return
	case '\n':
		s.line++
	case '<', '>', '=', '!':
		s.operators(r)
	case '{', '}', '(', ')', ',', ';', '.', '+', '-', '*':
		s.addToken(singleCharacters(r), string(r), nil)
	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(SLASH, "/", nil)
		}
	case '"':
		s.string()

	default:
		if utils.IsDigit(r) {
			s.number()
			return
		}
		if utils.IsAlpha(r) {
			s.identifier()
			return
		}
		s.error(fmt.Sprintf("Unexpected character: %c", r))
		s.setExitCode(65)

	}
}
func (s *Scanner) identifier() {
	for utils.IsAlphaNumeric(s.peek()) {
		s.advance()
	}

	text := string(s.source[s.start:s.current])
	t := matchKeywords(text)
	if t == "" {
		t = IDENTIFIER
	}

	s.addToken(t, text, nil)
}
func (s *Scanner) number() {
	for utils.IsDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == '.' && utils.IsDigit(s.peekNext()) {
		s.advance()

		for utils.IsDigit(s.peek()) {
			s.advance()
		}
	}

	value := string(s.source[s.start:s.current])
	fvalue, _ := strconv.ParseFloat(value, 64)

	s.addToken(NUMBER, value, fvalue)
}

func (s Scanner) error(message string) {
	s.report(s.line, "", message)
}

func (s Scanner) report(line int, where, message string) {
	fmt.Fprintf(os.Stderr, "[line %d] Error%s: %s\n", line, where, message)
}

func (s *Scanner) operators(r rune) {
	op := string(r)
	if s.match('=') {
		op += "="
	}

	s.addToken(matchOperators(op), op, nil)
}

func (s *Scanner) string() {
	value := strings.Builder{}
	for !s.isAtEnd() {
		ch := s.peek()
		if ch == '"' {
			s.advance()
			lexeme := string(s.source[s.start:s.current])
			s.addToken(STRING, lexeme, value.String())
			return
		}
		if ch == '\\' {
			s.advance()
			if s.isAtEnd() {
				s.error("Unterminated string.")
				return
			}
			escaped, _ := s.advance()
			switch escaped {
			case '"':
				value.WriteRune('"')
			case '\\':
				value.WriteRune('\\')
			case 'n':
				value.WriteRune('\n')
			case 't':
				value.WriteRune('\t')
			default:
				s.error(fmt.Sprintf("Unknown escape sequence: \\%c", escaped))
			}
			continue
		}
		if ch == '\n' {
			s.error("Unterminated string.")
			return
		}
		ch, _ = s.advance()
		value.WriteRune(ch)
	}
	s.error("Unterminated string.")
}

func (s *Scanner) match(r rune) bool {
	if s.current >= len(s.source) {
		return false
	}
	if rune(s.source[s.current]) != r {
		return false
	}
	s.current++
	return true
}

func (s *Scanner) addToken(tokenType TokenType, lexeme string, literal any) {
	s.tokens = append(s.tokens, Token{
		tokenType: tokenType,
		lexeme:    lexeme,
		literal:   literal,
		line:      s.line,
	})
}

func (s Scanner) GetTokens() []Token {
	return s.tokens
}

func (s Scanner) peek() rune {
	if s.isAtEnd() {
		return 0
	}
	return rune(s.source[s.current])
}
func (s Scanner) peekNext() rune {
	if s.current+1 >= len(s.source) {
		return 0
	}
	return rune(s.source[s.current+1])
}

func (s *Scanner) setExitCode(code int) {
	s.exitCode = code
}

// Tokens
