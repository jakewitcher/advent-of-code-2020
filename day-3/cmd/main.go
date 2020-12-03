package main

import (
	"day-3/internal/input"
	"day-3/internal/toboggan"
	"log"
)

func main() {
	matrix, err := input.Extract()
	if err != nil {
		log.Fatal(err.Error())
	}

	PartOne(matrix)
	PartTwo(matrix)
}

func PartOne(matrix [][]rune) {
	slope := toboggan.Slope{Right: 3, Down: 1}
	trees := toboggan.EvaluateSleddingPath(matrix, slope)

	log.Printf("trees along slope %v: %d", slope, trees)
}

func PartTwo(matrix [][]rune) {
	slopes := []toboggan.Slope{
		{Right: 1, Down: 1},
		{Right: 3, Down: 1},
		{Right: 5, Down: 1},
		{Right: 7, Down: 1},
		{Right: 1, Down: 2},
	}

	product := 1
	for _, slope := range slopes {
		trees := toboggan.EvaluateSleddingPath(matrix, slope)
		product *= trees
	}

	log.Printf("product of all trees encountered: %d", product)
}
