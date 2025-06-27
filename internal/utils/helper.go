package utils

import "strings"

func IsDigit(c rune) bool {
	return '0' <= c && c <= '9'
}

func IsAlpha(c rune) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || c == '_'
}

func IsAlphaNumeric(c rune) bool {
	return IsAlpha(c) || IsDigit(c)
}

func SnakeToTitle(snake string) string {
	words := strings.Split(snake, "_")
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToTitle(string(word[0])) + word[1:]
		}
	}
	return strings.Join(words, "")
}
