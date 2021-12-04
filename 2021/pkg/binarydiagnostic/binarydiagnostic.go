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
	sumArray := make([]int, inputLength)

	for _, input := range inputs {
		for i := 0; i < inputLength; i++ {
			if rune(input[i]) == rune('1') {
				sumArray[i] = sumArray[i] + 1
			}
		}
	}

	exponent := inputLength - 1
	for i := 0; i < inputLength; i++ {
		adder := int(math.Pow(2, float64(exponent)))
		if sumArray[i] > lenInputs/2 {
			gammaRate = gammaRate + adder
		} else {
			epsilonRate = epsilonRate + adder
		}

		exponent--
	}

	return fmt.Sprintf("%v", gammaRate*epsilonRate)
}

func (p *Puzzle) PartTwo() string {
	if len(inputs) <= 0 {
		log.Fatal()
	}

	oxygenGeneratorRating := oxygenFilter(inputs, 0)
	co2ScrubberRating := co2Filter(inputs, 0)

	oxygenRating := binaryStringToInt(oxygenGeneratorRating)
	co2Rating := binaryStringToInt(co2ScrubberRating)

	return fmt.Sprintf("%v", oxygenRating*co2Rating)
}

func oxygenFilter(arr []string, pos int) string {
	// end case
	if len(arr) == 1 {
		return arr[0]
	}
	// guard clauses
	if len(arr) <= 0 {
		log.Fatal()
	}
	if pos >= len(arr[0]) {
		log.Fatal()
	}

	onesCount := 0
	zerosCount := 0
	for _, v := range arr {
		if rune(v[pos]) == rune('1') {
			onesCount++
		} else {
			zerosCount++
		}
	}

	var mostCommonValue rune
	if onesCount >= zerosCount {
		mostCommonValue = rune('1')
	} else {
		mostCommonValue = rune('0')
	}

	filterArr := []string{}
	for _, v := range arr {
		if rune(v[pos]) == mostCommonValue {
			filterArr = append(filterArr, v)
		}
	}

	// recursion babyyyyyyyy
	return oxygenFilter(filterArr, pos+1)
}

func co2Filter(arr []string, pos int) string {
	// end case
	if len(arr) == 1 {
		return arr[0]
	}
	// guard clauses
	if len(arr) <= 0 {
		log.Fatal()
	}
	if pos >= len(arr[0]) {
		log.Fatal()
	}

	onesCount := 0
	zerosCount := 0
	for _, v := range arr {
		if rune(v[pos]) == rune('1') {
			onesCount++
		} else {
			zerosCount++
		}
	}

	var leastCommonValue rune
	if onesCount < zerosCount {
		leastCommonValue = rune('1')
	} else {
		leastCommonValue = rune('0')
	}

	filterArr := []string{}
	for _, v := range arr {
		if rune(v[pos]) == leastCommonValue {
			filterArr = append(filterArr, v)
		}
	}

	// recursion babyyyyyyyy
	return co2Filter(filterArr, pos+1)
}

func binaryStringToInt(b string) int {
	exponent := len(b) - 1
	result := 0
	for i := 0; i < len(b); i++ {
		if rune(b[i]) == rune('1') {
			adder := int(math.Pow(2, float64(exponent)))
			result = result + adder
		}
		exponent--
	}

	return result
}
