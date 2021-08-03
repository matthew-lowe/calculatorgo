package interpret

import (
	"unicode"
)

func Tokenize(source string) []*Token {
	tokens := make([]*Token, 0, len(source)) // Length of source is upper bound for the capacity of the slice

	// Convert the source to a list of tokens
	for _, char := range source {
		tokenType := findType(char)

		switch tokenType {
		case WHITESPACE:
		case NUM:
			value := int(char)
			tokens = append(tokens, NewToken(NUM, true, float64(value)-48)) // Subtract 48 becuase ascii
		default:
			tokens = append(tokens, NewToken(tokenType, false, 0))
		}
	}

	return mergeDigits(tokens)
}

func mergeDigits(tokens []*Token) []*Token {
	newTokens := make([]*Token, 0, len(tokens))
	newTokens = append(newTokens, NewToken(WHITESPACE, false, 0)) // ez solution to index errors

	for _, t := range tokens {
		if t.Type != NUM {
			newTokens = append(newTokens, t)
			continue
		}

		last := newTokens[len(newTokens)-1]

		if last.Type == NUM {
			oldValue := last.Value
			last.Value = oldValue*10 + t.Value
		} else {
			newTokens = append(newTokens, t)
		}
	}

	return newTokens[1:] // don't need the first whitespace character
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
	case '(':
		return L_BRACKET
	case ')':
		return R_BRACKET
	default:
		return NONE
	}
}

func removeWhitespace(tokens []*Token) []*Token {
	newTokens := make([]*Token, 0, len(tokens))

	for _, t := range tokens {
		if t.Type != WHITESPACE {
			newTokens = append(newTokens, t)
		}
	}

	return newTokens
}
