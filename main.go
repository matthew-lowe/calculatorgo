package main

import (
	"bufio"
	"fmt"
	"os"
)

type StringSerializable interface {
	ToString() string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter name: ")
	scanner.Scan()
	fmt.Printf("Hello %v!\n", scanner.Text())
}
