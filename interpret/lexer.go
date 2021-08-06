package interpret

import (
	"math"
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

	return mergeDecimals(mergeDigits(tokens))
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

func mergeDecimals(tokens []*Token) []*Token {
	newTokens := make([]*Token, 0, len(tokens))
	newTokens = append(newTokens, NewToken(WHITESPACE, false, 0))
	newTokens = append(newTokens, tokens[0])

	for i := 1; i < len(tokens); i++ {
		if i >= len(tokens)-1 || tokens[i].Type != DOT {
			newTokens = append(newTokens, tokens[i])
			continue
		} else {
			nextValue := tokens[i+1].Value
			prevValue := tokens[i-1].Value
			numDigits := math.Floor(math.Log10(nextValue) + 1)
			newTokens[len(newTokens)-1].Value = prevValue + (nextValue / math.Pow(10, numDigits))
			//fmt.Printf("new value: %v\n", )
			i += 1
		}
	}

	return newTokens[1:]
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
	case '.':
		return DOT
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
