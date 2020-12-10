package main

import (
	"day-10/internal/input"
	"day-10/internal/jolt"
	"log"
)

func main() {
	adapters, err := input.Extract()
	if err != nil {
		log.Fatal(err.Error())
	}

	PartOne(adapters)
	PartTwo(adapters)
}

func PartOne(adapters []int) {
	difference := jolt.CalculateDifference(adapters)
	log.Printf("product of adapters with difference of 1 and difference of 3: %d", difference)
}

func PartTwo(adapters []int) {
	paths := jolt.CalculatePaths(adapters)
	log.Printf("number of unique adapter paths: %d", paths)
}
