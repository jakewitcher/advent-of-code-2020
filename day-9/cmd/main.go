package main

import (
	"day-9/internal/input"
	"day-9/internal/xmas"
	"log"
)

func main() {
	numberSequence, err := input.Extract("internal/input/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	PartOne(numberSequence)
	PartTwo(numberSequence)
}

func PartOne(numberSequence []int) {
	n, err := xmas.FindFirstWeakness(numberSequence, 25)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("first encryption weakness in xmas sequence: %d", n)
}

func PartTwo(numberSequence []int) {
	n, err := xmas.FindSecondWeakness(numberSequence, 25)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("second encryption weakness in xmas sequence: %d", n)
}
