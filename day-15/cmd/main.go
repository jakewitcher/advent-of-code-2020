package main

import (
	"day-15/internal/game"
	"day-15/internal/input"
	"log"
)

func main() {
	start := input.Extract()
	PartOne(start)
	PartTwo(start)
}

func PartTwo(start []int) {
	nth := game.FindNthNumberSpoken(start, 30_000_000)
	log.Printf("2020th number spoken: %d", nth)
}

func PartOne(start []int) {
	nth := game.FindNthNumberSpoken(start, 2020)
	log.Printf("30_000_000th number spoken: %d", nth)
}