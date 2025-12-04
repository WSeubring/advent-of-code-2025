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
		parts := strings.Split(productRange, "-")
		first := parts[0]
		if !isValidProductCode(first) {
			number, err := strconv.Atoi(first)
			if err != nil {
				panic(err)
			}
			sum += number
		}
		second := parts[1]
		if !isValidProductCode(second) {
			number, err := strconv.Atoi(second)
			if err != nil {
				panic(err)
			}
			sum += number
		}
	}
	println("Sum of invalid product codes:", sum)
}

func isValidProductCode(code string) bool {
	if len(code)%2 == 1 {
		return false
	}

	// Split the string into two halves
	mid := len(code) / 2
	firstHalf := code[:mid]
	secondHalf := code[mid:]

	if firstHalf != secondHalf {
		return true
	}

	return false
}
