package scanner

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"chai-lang/internal/ast"
	"chai-lang/internal/errors"
	"chai-lang/internal/utils"
)

func singleCharacters(c rune) ast.TokenType {
	chars := map[rune]ast.TokenType{
		'{': ast.LEFT_BRACE,
		'}': ast.RIGHT_BRACE,
		'(': ast.LEFT_PAREN,
		')': ast.RIGHT_PAREN,
		',': ast.COMMA,
		';': ast.SEMICOLON,
		'.': ast.DOT,
		'+': ast.PLUS,
		'-': ast.MINUS,
		'*': ast.STAR,
		'/': ast.SLASH,
	}
	return chars[c]
}

func matchOperators(op string) ast.TokenType {
	operators := map[string]ast.TokenType{
		"==": ast.EQUAL_EQUAL,
		"<=": ast.LESS_EQUAL,
		">=": ast.GREATER_EQUAL,
		">":  ast.GREATER,
		"<":  ast.LESS,
		"!=": ast.BANG_EQUAL,
	}
	return operators[op]
}

func matchKeywords(s string) ast.TokenType {
	keywords := map[string]ast.TokenType{
		"hai":        ast.HAI,
		"chaibana":   ast.CHAIBANA,
		"chaikhatam": ast.CHAIKHATAM,
		"dekh":       ast.DEKH,
		"bol":        ast.BOL,
		"sun":        ast.SUN,
		"agar":       ast.AGAR,
		"nahito":     ast.NAHI_TO,
		"warna":      ast.WARNA,
		"jab_tak":    ast.JAB_TAK,
		"tab_tak":    ast.TAB_TAK,
		"har":        ast.HAR,
		"me":         ast.ME,
		"kaam":       ast.KAAM,
		"kar":        ast.KAR,
		"haan":       ast.HAAN,
		"nahi":       ast.NAHI,
		"khali":      ast.KHALI,
		"rehne_de":   ast.REHNE_DE,
		"yele":       ast.YE_LE,
		"ulta":       ast.ULTA,
	}

	return keywords[s]
}

type Scanner struct {
	source   []byte
	tokens   []ast.Token
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
			s.addToken(ast.SLASH, "/", nil)
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
		t = ast.IDENTIFIER
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

	s.addToken(ast.NUMBER, value, fvalue)
}

func (s Scanner) error(message string) {
	errors.ReportError(s.line, "", message)
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
			s.addToken(ast.STRING, lexeme, value.String())
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

func (s *Scanner) addToken(tokenType ast.TokenType, lexeme string, literal any) {
	s.tokens = append(s.tokens, ast.NewToken(tokenType, lexeme, literal, s.line))
}

func (s Scanner) GetTokens() []ast.Token {
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
