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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter maffs:")

	for {
		fmt.Print("> ")
		scanner.Scan()
		source := scanner.Text()
		tokens := interpret.Tokenize(source)
		printList(tokens)
		tree := interpret.Parse(tokens)
		fmt.Println(interpret.Evaluate(tree))
	}
}
