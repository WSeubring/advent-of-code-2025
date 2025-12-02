package day1

import (
	"aoc/utils"
	"fmt"
	"math"
	"strconv"
)

type Safe struct {
	CurrentPosition int
	ZeroHits        int
}

func NewSafe() *Safe {
	return &Safe{
		CurrentPosition: 50,
		ZeroHits:        0,
	}
}

func (s *Safe) Rotate(steps int) {
	start := s.CurrentPosition
	end := start + steps

	floorDiv := func(n int) int {
		return int(math.Floor(float64(n) / 100.0))
	}

	var hits int
	if steps > 0 {
		hits = floorDiv(end) - floorDiv(start)
	} else {
		hits = floorDiv(start-1) - floorDiv(end-1)
	}
	s.ZeroHits += hits

	s.CurrentPosition = end % 100
	if s.CurrentPosition < 0 {
		s.CurrentPosition += 100
	}
}

type Rotation struct {
	Steps     int
	Direction string
}

func (r *Rotation) ApplyToSafe(s *Safe) {
	s.Rotate(r.GetRotationSteps())
}

func (r *Rotation) GetRotationSteps() int {
	if r.Direction == "L" {
		return -r.Steps
	}
	return r.Steps
}

func RotationFromString(input string) (*Rotation, error) {
	direction := string(input[0])
	steps, err := strconv.Atoi(input[1:])
	if err != nil {
		return nil, err
	}

	return &Rotation{
		Steps:     steps,
		Direction: direction,
	}, nil
}

func SolvePart1AndPart2() {
	safe := NewSafe()

	lines, err := utils.ReadLines("day1/input.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	rotations := make([]*Rotation, 0, len(lines))
	for _, line := range lines {
		rotation, err := RotationFromString(line)
		if err != nil {
			fmt.Println("Error parsing rotation:", err)
			return
		}
		rotations = append(rotations, rotation)
	}

	zeroCount := 0
	for _, rotation := range rotations {
		rotation.ApplyToSafe(safe)

		if safe.CurrentPosition == 0 {
			zeroCount++
		}
	}

	fmt.Println("Part 1 (Zero at end):", zeroCount)
	fmt.Println("Part 2 (Total Zero hits):", safe.ZeroHits)
}
