package main

import (
	"day-2/internal/input"
	"day-2/internal/validate"
	"log"
)

func main() {
	passwords, err := input.Extract()
	if err != nil {
		log.Fatal(err.Error())
	}

	PartOne(passwords)
	PartTwo(passwords)
}

func PartOne(passwords []string) {
	valid, err := validate.IdentifyValidPasswords(passwords, validate.ApplySledRentalPolicy)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("valid sled rental passwords: %d", valid)
}

func PartTwo(passwords []string) {
	valid, err := validate.IdentifyValidPasswords(passwords, validate.ApplyTobogganCorpPolicy)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("valid toboggan corp passwords: %d", valid)
}
