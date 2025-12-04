package main

import (
	"aoc/day2"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	day2.SolvePart1()
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)
}
