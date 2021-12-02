package main

import (
	"log"

	"github.com/iweller-IQ/advent_of_code/2021/pkg/dive"
	"github.com/iweller-IQ/advent_of_code/2021/pkg/puzzle"
	"github.com/iweller-IQ/advent_of_code/2021/pkg/sonarsweep"
)

func main() {
	puzzles := map[string]puzzle.Puzzle{
		"Day 1": &sonarsweep.Puzzle{},
		"Day 2": &dive.Puzzle{},
	}

	for day, puzzle := range puzzles {
		log.Printf("%v.1 => %v", day, puzzle.PartOne())
		log.Printf("%v.2 => %v", day, puzzle.PartTwo())
	}
}
