package lanternfish

import "fmt"

type Puzzle struct{}

func (p *Puzzle) PartOne() string {
	lanternFish := [9]int{}
	days := 80

	// fill array
	for _, i := range inputs {
		lanternFish[i]++
	}

	// simulate fish
	buffer := [9]int{}
	for i := 0; i < days; i++ {
		buffer[8] = 0
		// shuffle the fish

		buffer[7] = lanternFish[8]
		buffer[6] = lanternFish[7]
		buffer[5] = lanternFish[6]
		buffer[4] = lanternFish[5]
		buffer[3] = lanternFish[4]
		buffer[2] = lanternFish[3]
		buffer[1] = lanternFish[2]
		buffer[0] = lanternFish[1]

		// handle spawn
		buffer[8] = lanternFish[0]
		buffer[6] = buffer[6] + lanternFish[0]

		lanternFish = buffer
	}

	// count fish
	result := 0
	for i := 0; i < len(lanternFish); i++ {
		result = result + lanternFish[i]
	}

	return fmt.Sprintf("%v", result)
}

func (p *Puzzle) PartTwo() string {
	lanternFish := [9]int{}
	days := 256

	// fill array
	for _, i := range inputs {
		lanternFish[i]++
	}

	// simulate fish
	buffer := [9]int{}
	for i := 0; i < days; i++ {
		buffer[8] = 0
		// shuffle the fish

		buffer[7] = lanternFish[8]
		buffer[6] = lanternFish[7]
		buffer[5] = lanternFish[6]
		buffer[4] = lanternFish[5]
		buffer[3] = lanternFish[4]
		buffer[2] = lanternFish[3]
		buffer[1] = lanternFish[2]
		buffer[0] = lanternFish[1]

		// handle spawn
		buffer[8] = lanternFish[0]
		buffer[6] = buffer[6] + lanternFish[0]

		lanternFish = buffer
	}

	// count fish
	result := 0
	for i := 0; i < len(lanternFish); i++ {
		result = result + lanternFish[i]
	}

	return fmt.Sprintf("%v", result)
}
