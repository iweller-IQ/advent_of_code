package main

import (
	"log"

	"github.com/iweller-IQ/advent_of_code/2021/pkg/binarydiagnostic"
	"github.com/iweller-IQ/advent_of_code/2021/pkg/dive"
	"github.com/iweller-IQ/advent_of_code/2021/pkg/giantsquid"
	"github.com/iweller-IQ/advent_of_code/2021/pkg/hydrothermal"
	"github.com/iweller-IQ/advent_of_code/2021/pkg/lanternfish"
	"github.com/iweller-IQ/advent_of_code/2021/pkg/puzzle"
	"github.com/iweller-IQ/advent_of_code/2021/pkg/sonarsweep"
	"github.com/iweller-IQ/advent_of_code/2021/pkg/whales"
)

func main() {
	puzzles := map[string]puzzle.Puzzle{
		"Day 1": &sonarsweep.Puzzle{},
		"Day 2": &dive.Puzzle{},
		"Day 3": &binarydiagnostic.Puzzle{},
		"Day 4": &giantsquid.Puzzle{},
		"Day 5": &hydrothermal.Puzzle{},
		"Day 6": &lanternfish.Puzzle{},
		"Day 7": &whales.Puzzle{},
	}

	for day, puzzle := range puzzles {
		log.Printf("%v.1 => %v", day, puzzle.PartOne())
		log.Printf("%v.2 => %v", day, puzzle.PartTwo())
	}
}
