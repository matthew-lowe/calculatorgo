package interpret

import "fmt"

type Node struct {
	Value    *Token
	Children [2]*Node
}

func NewNode(value *Token) *Node {
	return &Node{Value: value, Children: [2]*Node{nil, nil}}
}

func (node Node) String() string {
	return node.Value.String()
}

func (node Node) Print() {
	node.print("")
}

func (node Node) print(side string) {
	fmt.Println(side + ": " + node.Value.String())

	if node.Value.Type != NUM && node.Children[0] != nil && node.Children[1] != nil {
		node.Children[0].print(side + "L")
		node.Children[1].print(side + "R")
	}
}
