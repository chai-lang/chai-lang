package utils

func IsDigit(c rune) bool {
	return '0' <= c && c <= '9'
}

func IsAlpha(c rune) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || c == '_'
}
