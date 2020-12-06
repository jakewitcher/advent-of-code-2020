package main

import (
	"day-6/internal/form"
	"day-6/internal/input"
	"log"
)

func main() {
	groups, err := input.Extract()
	if err != nil {
		log.Fatal(err.Error())
	}

	PartOne(groups)
	PartTwo(groups)
}

func PartOne(groups [][]string) {
	var sum int
	for _, group := range groups {
		sum += form.EvaluateTotalYesResponsesByGroup(group)
	}

	log.Printf("sum of total yes responses: %d", sum)
}

func PartTwo(groups [][]string) {
	var sum int
	for _, group := range groups {
		sum += form.EvaluateConsensusYesResponsesByGroup(group)
	}

	log.Printf("sum of consensus yes responses: %d", sum)
}
