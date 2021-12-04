package binarydiagnostic

import (
	"fmt"
	"log"
	"math"
)

type Puzzle struct {
}

func (p *Puzzle) PartOne() string {
	gammaRate := 0
	epsilonRate := 0
	lenInputs := len(inputs)
	if lenInputs <= 0 {
		log.Fatal()
	}

	inputLength := len(inputs[0])
	sumMap := map[int]int{}

	for _, input := range inputs {
		for j := 0; j < inputLength; j++ {
			// ascii shenanigans
			sumMap[j] = sumMap[j] + int([]rune(input)[j]) - 48
		}
	}

	exponent := inputLength - 1
	for i := 0; i < inputLength; i++ {
		// ascii shenanigans
		if sumMap[i] > lenInputs/2 {
			gammaRate = gammaRate + int(math.Pow(2, float64(i)))
		}

		exponent--
	}

	epsilonRate = int(^byte(gammaRate))

	return fmt.Sprintf("%v", gammaRate*epsilonRate)
}

func (p *Puzzle) PartTwo() string {
	return "incomplete"
}
