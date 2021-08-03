package interpret

func getTokenPriority() [5]TokenType {
	return [5]TokenType{SUB, ADD, MUL, DIV, NUM}
}

func Parse(input []*Token) *Node {
	tokens := removeOuterBrackets(input)
	rootToken, rootIndex := getRoot(tokens)
	rootNode := NewNode(rootToken)

	// A number has no children, idiot!
	if rootNode.Value.Type == NUM {
		return rootNode
	}

	left := Parse(tokens[:rootIndex])
	right := Parse(tokens[rootIndex+1:])

	rootNode.Children[0] = left
	rootNode.Children[1] = right

	return rootNode
}

func getRoot(tokens []*Token) (*Token, int) {
	for _, p := range getTokenPriority() {
		for i, t := range tokens {
			if t.Type == p && !isInBrackets(tokens, i) {
				return t, i
			}
		}
	}

	return nil, 0 // It should never reach this, but stops compiler complaining
}

func isInBrackets(tokens []*Token, index int) bool {
	lCount := 0
	rCount := 0

	for _, token := range tokens[:index] {
		switch token.Type {
		case L_BRACKET:
			lCount++
		case R_BRACKET:
			rCount++
		}
	}

	return lCount != rCount
}

func shouldRemoveBrackets(tokens []*Token) bool {
	count := 0

	for index, token := range tokens {
		switch token.Type {
		case L_BRACKET:
			count--
		case R_BRACKET:
			count++
		}

		if count == 0 && index != (len(tokens)-1) {
			return false
		}
	}

	return true
}

func removeOuterBrackets(tokens []*Token) []*Token {
	if len(tokens) == 0 {
		return tokens
	}
	if !shouldRemoveBrackets(tokens) {
		return tokens
	}

	if tokens[0].Type == L_BRACKET && tokens[len(tokens)-1].Type == R_BRACKET {
		return tokens[1:(len(tokens) - 1)]
	}

	return tokens
}
