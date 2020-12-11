package main

import (
	"day-8/internal/game"
	"day-8/internal/input"
	"log"
)

func main() {
	instructions, err := input.Extract("internal/input/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	PartOne(instructions)
	PartTwo(instructions)
}

func PartOne(instructions []string) {
	accumulator, err := game.FindFinalAccumulatorBeforeLoop(instructions)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("final accumulator before loop: %d", accumulator)
}

func PartTwo(instructions []string) {
	accumulator, err := game.FindFinalAccumulatorOfFixedProgram(instructions)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("final accumulator after program is fixed: %d", accumulator)
}
