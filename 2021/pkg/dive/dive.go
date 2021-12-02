package dive

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Puzzle struct {
}

func (p *Puzzle) PartOne() string {
	horizontalPosition := 0
	depth := 0

	for _, input := range inputs {
		iArr := strings.Split(input, " ")
		if len(iArr) != 2 {
			log.Fatal()
		}

		direction := iArr[0]
		distance, _ := strconv.Atoi(iArr[1])
		switch direction {
		case "forward":
			distance = distance * 1
			horizontalPosition = horizontalPosition + distance
		case "up":
			distance = distance * -1
			depth = depth + distance
		case "down":
			distance = distance * 1
			depth = depth + distance
		}
	}

	result := horizontalPosition * depth

	return fmt.Sprint(result)
}

func (p *Puzzle) PartTwo() string {
	horizontalPosition := 0
	depth := 0
	aim := 0

	for _, input := range inputs {
		iArr := strings.Split(input, " ")
		if len(iArr) != 2 {
			log.Fatal()
		}

		direction := iArr[0]
		distance, _ := strconv.Atoi(iArr[1])
		switch direction {
		case "forward":
			distance = distance * 1
			horizontalPosition = horizontalPosition + distance
			depth = depth + (aim * distance)
		case "up":
			distance = distance * -1
			aim = aim + distance
		case "down":
			distance = distance * 1
			aim = aim + distance
		}
	}

	result := horizontalPosition * depth

	return fmt.Sprint(result)
}
