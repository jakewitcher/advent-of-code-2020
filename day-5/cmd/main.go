package main

import (
	"day-5/internal/boardingpass"
	"day-5/internal/input"
	"log"
	"sort"
)

func main() {
	boardingPasses, err := input.Extract()
	if err != nil {
		log.Fatal(err.Error())
	}

	PartOne(boardingPasses)
	PartTwo(boardingPasses)
}

func PartOne(boardingPasses []string) {
	var max int
	for _, b := range boardingPasses {
		boardingPass := boardingpass.Parse(b)
		if boardingPass.SeatId > max {
			max = boardingPass.SeatId
		}
	}

	log.Printf("max seat id number: %d", max)
}

func PartTwo(boardingPasses []string) {
	seatIds := make([]int, len(boardingPasses))
	for i, b := range boardingPasses {
		boardingPass := boardingpass.Parse(b)
		seatIds[i] = boardingPass.SeatId
	}

	sort.Ints(seatIds)

	prev := seatIds[0]
	for _, seatId := range seatIds {
		if prev+2 == seatId {
			log.Printf("seat id is %d", seatId-1)
			break
		}

		prev = seatId
	}
}
