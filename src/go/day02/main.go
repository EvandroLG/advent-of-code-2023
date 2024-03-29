package main

import (
	"fmt"
	"os"
)

func main() {
	input, err := os.ReadFile("./input.txt")

	if err != nil {
		panic(err)
	}

	firstResult := firstPart(string(input))
	fmt.Printf("Result 1: %d", firstResult)
	fmt.Print("\n")

	secondResult := secondPart(string(input))
	fmt.Printf("Result 2: %d", secondResult)
}
