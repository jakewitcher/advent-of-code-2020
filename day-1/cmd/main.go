package main

import (
	"day-1/internal/input"
	"day-1/internal/report"
	"log"
)

func main() {
	entries, err := input.Extract()
	if err != nil {
		log.Fatal(err.Error())
	}

	PartOne(entries)
	PartTwo(entries)
}

func PartOne(entries []int) {
	product, err := report.FindProductOfTwoEntriesWithSum(entries, 2020)
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Printf("product of entries with sum of 2020: %d", product)
}

func PartTwo(entries []int) {
	product, err := report.FindProductOfThreeEntriesWithSum(entries, 2020)
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Printf("product of entries with sum of 2020: %d", product)
}