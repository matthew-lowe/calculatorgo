package main

import (
	"testing"
)

/*
func BenchmarkCalculateStream(b *testing.B) {
	inputStream := make(chan string)
	tokenStream := interpret.TokenizeStream(inputStream)
	treeStream := interpret.ParseStream(tokenStream)
	outputStream := interpret.EvaluateStream(treeStream)

	go func() {
		for i := 0; i < b.N; i++ {
			inputStream <- "51235 + 1255235 - 1241525 * (131543 ^ 2) / 5"
		}
		close(inputStream)
	}()

	// Block until we're all funky
	for range outputStream {
	}
}
*/

func BenchmarkCalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		run("51235 + 1255235 - 1241525 * (131543 ^ 2) / 5")

	}
}
