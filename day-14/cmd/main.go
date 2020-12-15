package main

import (
	"day-14/internal/input"
	"day-14/internal/program"
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
	sum, err := program.RunInitialization(instructions, &program.VersionOneDecoder{})
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("v1 decoder, sum of values in memory: %d", sum)
}

func PartTwo(instructions []string) {
	sum, err := program.RunInitialization(instructions, &program.VersionTwoDecoder{})
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("v2 decoder, sum of values in memory: %d", sum)
}
