package main

import (
	"day-7/internal/input"
	"day-7/internal/luggage"
	"log"
)

func main() {
	rules, err := input.Extract()
	if err != nil {
		log.Fatal(err.Error())
	}

	PartOne(rules)
	PartTwo(rules)
}

func PartOne(rules []string) {
	count, err := luggage.SumPossibleOuterBags(rules, "shiny gold")
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("number of outer bags containing target: %d", count)
}

func PartTwo(rules []string) {
	count, err := luggage.SumRequiredInnerBags(rules, "shiny gold")
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("number of inner bags required by target: %d", count)
}
