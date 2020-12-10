package main

import (
	"day-10/internal/adapters"
	io "day-10/internal/input"
	"log"
)

func main() {
	input, err := io.Extract("internal/input/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	PartOne(input)
	PartTwo(input)
}

func PartOne(input []int) {
	difference := adapters.CalculateDifference(input)
	log.Printf("product of adapters with difference of 1 and difference of 3: %d", difference)
}

func PartTwo(input []int) {
	paths := adapters.CalculatePaths(input)
	log.Printf("number of unique adapter paths: %d", paths)
}
