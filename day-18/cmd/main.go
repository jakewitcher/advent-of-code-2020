package main

import (
	"day-18/internal/homework"
	"day-18/internal/input"
	"log"
)

func main() {
	expressions, err := input.Extract("internal/input/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	PartOne(expressions)
	PartTwo(expressions)
}

func PartOne(expressions []string) {
	sum, err := homework.SumEvaluatedExpressions(expressions, homework.EqualPrecedence{})
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("sum of all evaluated expressions: %d", sum)
}

func PartTwo(expressions []string) {
	sum, err := homework.SumEvaluatedExpressions(expressions, homework.AdditionHigherPrecedence{})
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("sum of all evaluated expressions: %d", sum)
}
