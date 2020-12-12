package main

import (
	"day-11/internal/input"
	"day-11/internal/seating"
	"log"
)

func main() {
	matrix, err := input.Extract("internal/input/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	PartOne(matrix)
	PartTwo(matrix)
}

func PartOne(matrix [][]rune) {
	occupied, err := seating.Normalize(matrix, seating.AdjacentRules{})
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("number of occupied seats in normalized seating arrangement: %d", occupied)
}

func PartTwo(matrix [][]rune) {
	occupied, err := seating.Normalize(matrix, seating.VisibleRules{})
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("number of occupied seats in normalized seating arrangement: %d", occupied)
}
