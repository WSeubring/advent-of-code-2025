package day4

import (
	"aoc/utils"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func Solve() {
	fmt.Println("Solving Day 4")
	lines := strings.Split(input, "\n")
	grid := utils.StringsToRunes(lines)
	factoryFloor := NewFactoryFloor(grid)

	count := 0
	noRemoves := true
	for {
		paperRolls := factoryFloor.GetAllPaperRollTiles()
		for _, roll := range paperRolls {
			if factoryFloor.CanRollOfPaperBePickedUp(roll) {
				factoryFloor.RemoveRollOfPaper(roll)
				noRemoves = false
				count++
			}
		}
		if noRemoves {
			break
		}

		noRemoves = true
	}

	fmt.Printf("Number of paper rolls that can be picked up: %d\n", count)

}

func (ff FactoryFloor) RemoveRollOfPaper(tile FloorTile) {
	tile.Rune = '.'
	ff[tile.Position.Y][tile.Position.X] = tile
}

func (ff FactoryFloor) CanRollOfPaperBePickedUp(tile FloorTile) bool {
	return ff.NumberAdjacentRollsOfPaper(tile.Position) < 4
}

func (ff FactoryFloor) GetAllPaperRollTiles() []FloorTile {
	var paperRolls []FloorTile
	for _, row := range ff {
		for _, tile := range row {
			if tile.IsRollOfPaper() {
				paperRolls = append(paperRolls, tile)
			}
		}
	}
	return paperRolls
}

func (ff FactoryFloor) NumberAdjacentRollsOfPaper(pos Position) int {
	directions := []Position{
		{X: -1, Y: -1}, {X: 0, Y: -1}, {X: 1, Y: -1},
		{X: -1, Y: 0} /*{X:0,Y:0},*/, {X: 1, Y: 0},
		{X: -1, Y: 1}, {X: 0, Y: 1}, {X: 1, Y: 1},
	}

	count := 0
	for _, dir := range directions {
		newX := pos.X + dir.X
		newY := pos.Y + dir.Y

		if newY >= 0 && newY < len(ff) && newX >= 0 && newX < len(ff[newY]) {
			if ff[newY][newX].IsRollOfPaper() {
				count++
			}
		}
	}
	return count
}

func NewFactoryFloor(runes [][]rune) FactoryFloor {
	floor := make(FactoryFloor, len(runes))
	for y, row := range runes {
		floor[y] = make([]FloorTile, len(row))
		for x, r := range row {
			floor[y][x] = FloorTile{
				Position: Position{X: x, Y: y},
				Rune:     r,
			}
		}
	}
	return floor
}

type Position struct {
	X int
	Y int
}

type FactoryFloor [][]FloorTile

type FloorTile struct {
	Position
	Rune rune
}

func (f FloorTile) IsRollOfPaper() bool {
	return f.Rune == '@'
}
