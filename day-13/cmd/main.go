package main

import (
	"day-13/internal/input"
	"day-13/internal/shuttle"
	"log"
)

func main() {
	timestamp, busIds, err := input.Extract("internal/input/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	PartOne(timestamp, busIds)
	PartTwo(busIds)
}

func PartOne(timestamp int, buses []string) {
	arrival, err := shuttle.FindNextAvailableBus(timestamp, buses)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("product of time until next bus arrives and bus id: %d", arrival)
}

func PartTwo(buses []string) {
	arrival, err := shuttle.FindEarliestTimestamp(buses)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("earliest timestamp meeting condition: %d", arrival)
}
