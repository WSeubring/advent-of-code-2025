package day5

import (
	"aoc/utils"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type FreshIngredientRange struct {
	Start int
	End   int
}

func InputToFreshIngredientRange(input string) FreshIngredientRange {
	fmt.Println("Parsing fresh ingredient range from input:", input)
	parts := strings.Split(input, "-")
	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])

	return FreshIngredientRange{Start: start, End: end}
}

func (fir FreshIngredientRange) CanMerge(other FreshIngredientRange) bool {
	return fir.Start <= other.End+1 && other.Start <= fir.End+1
}

func (fir FreshIngredientRange) Merge(other FreshIngredientRange) FreshIngredientRange {
	fmt.Println("Merging ranges:", fir, "and", other)

	start := fir.Start
	if other.Start < start {
		start = other.Start
	}
	end := fir.End
	if other.End > end {
		end = other.End
	}

	fmt.Println("Merged range:", FreshIngredientRange{Start: start, End: end})
	return FreshIngredientRange{Start: start, End: end}
}

func ReduceRanges(freshRanges []FreshIngredientRange) []FreshIngredientRange {
	i := 0
	for i < len(freshRanges) {
		current := freshRanges[i]
		fmt.Println("Checking range for merges:", current)
		j := i + 1
		merged := false
		for j < len(freshRanges) {
			other := freshRanges[j]

			if current.CanMerge(other) {
				freshRanges[i] = current.Merge(other)
				fmt.Println("Removing merged range at index", j, ":", freshRanges[j])
				freshRanges = utils.RemoveIndexFromArr(freshRanges, j)
				current = freshRanges[i]
				merged = true
			} else {
				j++
			}
		}
		if !merged {
			i++
		}
	}
	return freshRanges
}

func (fir FreshIngredientRange) IsProductInRange(product int) bool {
	return product >= fir.Start && product <= fir.End
}

func (fir FreshIngredientRange) Range() int {
	// Inclusive range
	return fir.End - fir.Start + 1
}

func Solve() {
	lines := strings.Split(input, "\n")

	var freshRanges []FreshIngredientRange

	startingIndexProducts := 0
	for i, line := range lines {
		if line == "" {
			startingIndexProducts = i + 1
			break
		}
		freshRanges = append(freshRanges, InputToFreshIngredientRange(line))
	}

	for _, line := range lines[startingIndexProducts:] {
		fmt.Println(line)
	}

	freshRanges = ReduceRanges(freshRanges)
	for i, fr := range freshRanges {
		fmt.Printf("Fresh range %d: %d-%d\n", i, fr.Start, fr.End)
	}

	count := 0
	for _, line := range lines[startingIndexProducts:] {
		product, _ := strconv.Atoi(line)
		for _, freshRange := range freshRanges {
			if freshRange.IsProductInRange(product) {
				count++
				break
			}
		}
	}
	fmt.Println("Count of fresh products:", count)

	totalRange := 0
	for _, freshRange := range freshRanges {
		totalRange += freshRange.Range()
	}
	fmt.Println("Total fresh ingredient range:", totalRange)
}
