package sonarsweep

import "fmt"

type Puzzle struct {
}

func (p *Puzzle) PartOne() string {
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

	return fmt.Sprintf("%v", result)
}

func (p *Puzzle) PartTwo() string {
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

	return fmt.Sprintf("%v", result)
}
