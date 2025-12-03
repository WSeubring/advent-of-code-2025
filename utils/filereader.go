package utils

import (
	"bufio"
	"os"
)

func ReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func ParseStringOfNumbersToInts(input string) []int {
	var numbers []int
	for _, char := range input {
		numbers = append(numbers, CastCharToInt(char))
	}
	return numbers
}

func CastCharToInt(c rune) int {
	return int(c - '0')
}

func GetIndexOfMaxInt(ints []int) int {
	maxIndex := 0
	maxValue := ints[0]

	for i, v := range ints {
		if v > maxValue {
			maxValue = v
			maxIndex = i
		}
	}
	return maxIndex
}
