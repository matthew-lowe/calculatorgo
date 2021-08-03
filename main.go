package main

import (
	"fmt"

	"github.com/matthewlowe/calculatorgo/interpret"
)

func printList(tokens []*interpret.Token) {
	for _, t := range tokens {
		fmt.Printf("%v, ", (*t).String())
	}
	fmt.Println()
}

func main() {
	/*
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter name: ")
		scanner.Scan()
		fmt.Printf("Hello %v!\n", scanner.Text())
	*/

	source := "(69 + 420) * (666 / 2)"
	tokens := interpret.Tokenize(source)
	tree := interpret.Parse(tokens)
	fmt.Println(interpret.Evaluate(tree))

}
