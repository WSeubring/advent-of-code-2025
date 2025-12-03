package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimple12SizeInputPart2(t *testing.T) {
	input := "123456789012"
	batteryBank := NewBatteryBank(input)

	maxVoltage := batteryBank.GetMaximumVoltagePart2()

	assert.Equal(t, 123456789012, maxVoltage)
}

func TestLargerInputPart2(t *testing.T) {
	input := "9811111111111111111111116"
	batteryBank := NewBatteryBank(input)

	maxVoltage := batteryBank.GetMaximumVoltagePart2()

	assert.Equal(t, 981111111116, maxVoltage)
}

func TestComplexInputPart2(t *testing.T) {
	input := "981111111111118111111111111111119111111111111111111111111811111111111111111112"
	batteryBank := NewBatteryBank(input)

	maxVoltage := batteryBank.GetMaximumVoltagePart2()

	assert.Equal(t, 998111111112, maxVoltage)
}
