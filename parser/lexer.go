package parser

import (
	"strings"
)

type Token struct {
	Type  string
	Value string
}

func Lex(input string) []Token {
	tokens := []Token{}
	words := strings.Fields(strings.ToUpper(input))

	for _, word := range words {
		switch word {
		case "SELECT", "FROM", "INSERT", "INTO", "VALUES", "CREATE", "TABLE":
			tokens = append(tokens, Token{Type: "KEYWORD", Value: word})
		default:
			tokens = append(tokens, Token{Type: "IDENTIFIER", Value: word})
		}
	}

	return tokens
}
