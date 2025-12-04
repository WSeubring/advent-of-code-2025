package day2

import (
	"aoc/utils"
	"strconv"
	"strings"
)

func SolvePart1() {
	data, err := utils.GetFileAsString("day2/input.txt")
	if err != nil {
		panic(err)
	}

	data = strings.ReplaceAll(data, "\n", "")

	sum := 0

	productRanges := strings.Split(data, ",")
	for _, productRange := range productRanges {
		pr := InputToProductRange(productRange)
		invalidCodes := pr.GetInvalidCodes()
		for _, code := range invalidCodes {
			sum += code
		}
	}

	println("Sum of invalid product codes:", sum)
}

type ProductRange struct {
	Start int
	End   int
}

func InputToProductRange(input string) ProductRange {
	parts := strings.Split(input, "-")
	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])
	return ProductRange{Start: start, End: end}
}

func (pr ProductRange) GetInvalidCodes() []int {
	var invalidCodes []int
	for i := pr.Start; i <= pr.End; i++ {
		codeStr := CodeAsString(i)
		if !isValidProductCode(codeStr) {
			invalidCodes = append(invalidCodes, i)
		}
	}
	return invalidCodes
}

// Look at me wasting cpu cycles...
func (pr ProductRange) ExplodeRangeToFullList() []int {
	var codes []int
	for i := pr.Start; i <= pr.End; i++ {
		codes = append(codes, i)
	}
	return codes
}

func CodeAsString(code int) string {
	return strconv.Itoa(code)
}

// func isValidProductCodePart1(code string) bool {
// 	if len(code)%2 == 1 {
// 		return true
// 	}

// 	// Split the string into two halves
// 	mid := len(code) / 2
// 	firstHalf := code[:mid]
// 	secondHalf := code[mid:]

// 	return firstHalf != secondHalf
// }

func isValidProductCode(code string) bool {
	for i := 0; i < len(code)/2; i += 1 {
		seq := code[:i+1]
		expectedRepetitions := len(code) / len(seq)
		reconstructed := strings.Repeat(seq, expectedRepetitions)
		if reconstructed == code {
			return false
		}
	}

	return true
}
