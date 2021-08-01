package interpret

import "fmt"

func getTokenPriority() [5]TokenType {
	return [5]TokenType{DIV, MUL, ADD, SUB, NUM}
}

func Parse(tokens []*Token) *Node {
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
			if t.Type == p {
				return t, i
			}
		}
	}
	fmt.Println("goujon!!!!!!!!!!!!!")

	return nil, 0 // It should never reach this, but stops compiler complaining
}
