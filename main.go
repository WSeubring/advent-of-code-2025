package main

import (
	"aoc/day5"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	day5.Solve()

	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)
}
