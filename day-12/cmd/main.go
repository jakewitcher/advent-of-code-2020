package main

import (
	"day-12/internal/input"
	"day-12/internal/navigation"
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
	distance, err := navigation.FindManhattanDistance(instructions, navigation.NaiveInterpreter{})
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("naive interpreter manhattan distance: %d", distance)
}

func PartTwo(instructions []string) {
	distance, err := navigation.FindManhattanDistance(instructions, navigation.WaypointInterpreter{})
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("waypoint interpreter manhattan distance: %d", distance)
}
