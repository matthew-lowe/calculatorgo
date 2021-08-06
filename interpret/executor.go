package interpret

import "math"

func EvaluateStream(in <-chan *Node) <-chan float64 {
	out := make(chan float64)
	go func() {
		for n := range in {
			out <- Evaluate(n)
		}
		close(out)
	}()
	return out
}

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
	case EXP:
		return math.Pow(Evaluate(root.Children[0]), Evaluate(root.Children[1]))
	default:
		panic("FUCK")
	}
}
