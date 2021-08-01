package interpret

func Evaluate(root *Node) float64 {
	switch root.Value.Type {
	case NUM:
		return root.Value.Value
	case ADD:
		return Evaluate(root.Children[0]) + Evaluate(root.Children[1])
	case SUB:
		return Evaluate(root.Children[0]) - Evaluate(root.Children[1])
	case MUL:
		return Evaluate(root.Children[0]) * Evaluate(root.Children[1])
	case DIV:
		return Evaluate(root.Children[0]) / Evaluate(root.Children[1])
	default:
		panic("FUCK")
	}
}
