package main

import "unicode"

func Tokenize(source string) []*Token {
	tokens := make([]*Token, 0, len(source)) // Length of source is upper bound for the capacity of the slice

	// Convert the source to a list of tokens
	for _, char := range source {
		tokenType := findType(char)

		if tokenType == NUM {
			value := int(char)
			tokens = append(tokens, NewToken(NUM, true, float64(value)))
		} else {
			tokens = append(tokens, NewToken(tokenType, false, 0))
		}
	}

	return tokens
}

func findType(char rune) TokenType {
	if unicode.IsDigit(char) {
		return NUM
	}

	switch char {
	case ' ':
		return WHITESPACE
	case '+':
		return ADD
	case '-':
		return SUB
	case '*':
		return MUL
	case '/':
		return DIV
	default:
		return NONE
	}
}
