package main

import (
	"day-4/internal/input"
	"day-4/internal/validate"
	"log"
)

func main() {
	passports, err := input.Extract()
	if err != nil {
		log.Fatal(err.Error())
	}

	valid, err := validate.AllPassports(passports)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("number of valid passports: %d", valid)
}
