package day3

import (
	"aoc/utils"
	"fmt"
	"strconv"
)

type BatteryBank struct {
	Batteries []int
}

func NewBatteryBank(input string) *BatteryBank {
	batteries := utils.ParseStringOfNumbersToInts(input)

	return &BatteryBank{
		Batteries: batteries,
	}
}

func (bb *BatteryBank) GetMaximumVoltage() int {
	firstIndexOfHighestBattery := utils.GetIndexOfMaxInt(bb.Batteries[:len(bb.Batteries)-1])
	firstBatteryVoltageDigit := bb.Batteries[firstIndexOfHighestBattery]

	remainingBatteries := bb.Batteries[firstIndexOfHighestBattery+1:]
	secondIndexOfHighestBattery := utils.GetIndexOfMaxInt(remainingBatteries)
	secondBatteryVoltageDigit := remainingBatteries[secondIndexOfHighestBattery]

	value, err := strconv.Atoi(strconv.Itoa(firstBatteryVoltageDigit) + strconv.Itoa(secondBatteryVoltageDigit))
	if err != nil {
		panic(err)
	}

	return value
}

func (bb *BatteryBank) GetMaximumVoltagePart2() int {
	numberOfBatteriesToSelect := 12
	selectedBatteries := make([]int, 0, numberOfBatteriesToSelect)
	remainingBatteries := bb.Batteries
	for i := 0; i < numberOfBatteriesToSelect; i++ {
		indexOfHighestBattery := utils.GetIndexOfMaxInt(remainingBatteries[:len(remainingBatteries)-(numberOfBatteriesToSelect-i-1)])
		selectedBatteries = append(selectedBatteries, remainingBatteries[indexOfHighestBattery])
		remainingBatteries = remainingBatteries[indexOfHighestBattery+1:]
	}

	valueStr := ""
	for _, digit := range selectedBatteries {
		valueStr += strconv.Itoa(digit)
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		panic(err)
	}

	return value
}

func SolvePart1() {
	bankLines, err := utils.ReadLines("day3/input.txt")
	if err != nil {
		panic(err)
	}

	sum := 0
	for _, line := range bankLines {
		batteryBank := NewBatteryBank(line)
		voltage := batteryBank.GetMaximumVoltage()

		sum += voltage
	}

	fmt.Println("Total sum of maximum voltages:", sum)
}

func SolvePart2() {
	bankLines, err := utils.ReadLines("day3/input.txt")
	if err != nil {
		panic(err)
	}

	sum := 0
	for _, line := range bankLines {
		batteryBank := NewBatteryBank(line)
		voltage := batteryBank.GetMaximumVoltagePart2()
		sum += voltage
	}

	fmt.Println("Total sum of maximum voltages (part 2):", sum)
}
