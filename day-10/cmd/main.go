package main

import (
	"day-10/internal/adapters"
	"day-10/internal/input"
	"log"
)

func main() {
	joltAdapters, err := input.Extract("internal/input/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	PartOne(joltAdapters)
	PartTwo(joltAdapters)
}

func PartOne(joltAdapters []int) {
	difference := adapters.CalculateDifference(joltAdapters)
	log.Printf("product of adapters with difference of 1 and difference of 3: %d", difference)
}

func PartTwo(joltAdapters []int) {
	paths := adapters.CalculatePaths(joltAdapters)
	log.Printf("number of unique adapter paths: %d", paths)
}
