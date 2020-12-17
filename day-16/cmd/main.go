package main

import (
	"day-16/internal/input"
	"day-16/internal/ticket"
	"log"
)

func main() {
	fields, yourTicket, nearbyTickets, err := input.Extract("internal/input/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	PartOne(fields, nearbyTickets)
	PartTwo(fields, yourTicket, nearbyTickets)
}

func PartOne(fields []string, tickets []string) {
	errorRate, err := ticket.ScanErrorRate(fields, tickets)
	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Printf("error rate: %d", errorRate)
}

func PartTwo(fields []string, yourTicket string, nearbyTickets []string) {
	departureFieldsProduct, err := ticket.IdentifyDepartureFields(fields, yourTicket, nearbyTickets)
	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Printf("product of departure tickets: %d", departureFieldsProduct)
}
