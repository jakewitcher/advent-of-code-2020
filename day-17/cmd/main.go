package main

import (
	"day-17/internal/dimension3D"
	"day-17/internal/dimension4D"
	"day-17/internal/input"
	"log"
)

func main() {
	cubes, err := input.Extract("internal/input/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	PartOne(cubes)
	PartTwo(cubes)
}

func PartOne(cubes [][]rune) {
	active, err := dimension3D.FindActiveCubesAfterNCycles(cubes, 6)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("number of active cubes: %d", active)
}

func PartTwo(cubes [][]rune) {
	active, err := dimension4D.FindActiveCubesAfterNCycles(cubes, 6)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("number of active cubes: %d", active)
}
