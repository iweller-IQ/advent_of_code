package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	inputs := strings.Split(string(content), "\n")
	inputArray := intArray(inputs)

	partOneAnswer := PartOne(inputArray)
	fmt.Printf("The answer to part one is %v\n", partOneAnswer)

	partTwoAnswer := PartTwo(inputArray)
	fmt.Printf("The answer to part two is %v\n", partTwoAnswer)
}

func PartOne(inputs []int) int {
	result := 0
	previousNumber := -1
	for i, number := range inputs {
		if i == 0 {
			previousNumber = number
			continue
		}

		if previousNumber < number {
			result++
		}

		previousNumber = number
	}

	return result
}

func PartTwo(inputs []int) int {
	result := 0
	arrayLen := len(inputs)

	previousSum := 0
	for i := range inputs {
		if i+2 >= arrayLen {
			break
		}

		if i == 0 {
			previousSum = inputs[i] + inputs[i+1] + inputs[i+2]
			continue
		}

		currentSum := inputs[i] + inputs[i+1] + inputs[i+2]
		if previousSum < currentSum {
			result++
		}

		previousSum = currentSum
	}

	return result
}

func intArray(sArr []string) []int {
	result := []int{}
	for _, s := range sArr {
		i, _ := strconv.Atoi(s)
		result = append(result, i)
	}

	return result
}
