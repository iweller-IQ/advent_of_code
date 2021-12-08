package whales

import (
	"fmt"
	"math"
)

var MaxInt = int(^uint(0) >> 1)

type Puzzle struct {
}

func (p *Puzzle) PartOne() string {
	minFuelUsed := MaxInt
	for _, guess := range inputs {
		fuelUsed := 0
		for _, input := range inputs {
			fuelUsed += int(math.Abs(float64(guess - input)))
		}

		if fuelUsed < minFuelUsed {
			minFuelUsed = fuelUsed
		}
	}

	return fmt.Sprintf("%v", minFuelUsed)
}

func (p *Puzzle) PartTwo() string {
	minFuelUsed := MaxInt
	minInput := MaxInt
	maxInput := 0
	for _, input := range inputs {
		if input > maxInput {
			maxInput = input
		}

		if input < minInput {
			minInput = input
		}
	}
	for guess := minInput; guess <= maxInput; guess++ {
		fuelUsed := 0
		for _, input := range inputs {
			deltaFuel := guess - input
			if deltaFuel < 0 {
				deltaFuel = deltaFuel * -1
			}
			fuelUsed += (deltaFuel * (deltaFuel + 1)) / 2
		}

		if fuelUsed < minFuelUsed {
			minFuelUsed = fuelUsed
		}
	}

	return fmt.Sprintf("%v", minFuelUsed)
}
