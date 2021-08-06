package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/matthewlowe/calculatorgo/interpret"
)

func printList(tokens []*interpret.Token) {
	for _, t := range tokens {
		fmt.Printf("%v ", (*t).String())
	}
	fmt.Println()
}

func run(source string) float64 {
	tokens := interpret.Tokenize(source)
	tree := interpret.Parse(tokens)
	return interpret.Evaluate(tree)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter maffs:")

	for {
		fmt.Print("> ")
		scanner.Scan()
		source := scanner.Text()
		fmt.Println(run(source))
	}
}
