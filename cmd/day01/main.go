package main

import (
	"fmt"
	"github.com/donavynhaley/advent-of-code-2023/pkg/day01"
	"log"
	"os"
	"path/filepath"
)

func main() {
	input, err := os.ReadFile(filepath.Join("..", "..", "input", "day01.txt"))
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}

	result := day01.SolvePartOne(string(input))
	//fmt.Printf("Part 1 Input: %v\n", string(input))
	fmt.Printf("Part 1 Result: %v\n", result)

	//result = day01.SolvePartTwo(string(input))
	//fmt.Printf("Part 2 Input: %v\n", string(input))
	//fmt.Printf("Part 2 Result: %v\n", result)
}
