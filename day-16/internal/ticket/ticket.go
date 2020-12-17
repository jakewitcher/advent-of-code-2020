package ticket

import (
	"strings"
)

type Set map[int]interface{}

func (s Set) Union(other Set) {
	for val := range other {
		s[val] = struct{}{}
	}
}

func (s Set) Contains(n int) bool {
	_, ok := s[n]
	return ok
}

type Field struct {
	Name           string
	ValidPositions []int
	TotalValid     int
	Set
}

type Ticket []int

func IdentifyDepartureFields(fieldsInput []string, yourTicketInput string, nearbyTicketsInput []string) (int, error) {
	fields, err := ParseFields(fieldsInput)
	if err != nil {
		return 0, err
	}

	yourTicket, err := ParseTicket(yourTicketInput)
	if err != nil {
		return 0, err
	}

	tickets, err := ParseTickets(nearbyTicketsInput)
	if err != nil {
		return 0, err
	}
	tickets = append(tickets, yourTicket)

	identifier := NewPositionIdentifier(fields)
	identifier.accumulateValidPositionsByField(tickets)
	identifier.sortFieldsByNumberOfValidPositions()

	fieldPositions := identifier.findActualPositionByField()

	departureFieldsProduct := 1
	for fieldName, position := range fieldPositions {
		if strings.HasPrefix(fieldName, "departure") {
			departureFieldsProduct *= yourTicket[position]
		}
	}

	return departureFieldsProduct, nil
}

func ScanErrorRate(fieldsInput []string, ticketsInput []string) (int, error) {
	var errorRate int
	fields, err := ParseFields(fieldsInput)
	if err != nil {
		return errorRate, err
	}

	tickets, err := ParseTickets(ticketsInput)
	if err != nil {
		return errorRate, err
	}

	invalidFieldValues := getInvalidFieldValues(fields, tickets)

	for _, value := range invalidFieldValues {
		errorRate += value
	}

	return errorRate, nil
}

func getInvalidFieldValues(fields []*Field, tickets []Ticket) []int {
	validFieldValues := getValidFieldValues(fields)

	invalidFieldValues := make([]int, 0)
	for _, ticket := range tickets {
		for _, value := range ticket {
			if !validFieldValues.Contains(value) {
				invalidFieldValues = append(invalidFieldValues, value)
			}
		}
	}
	return invalidFieldValues
}

func getValidFieldValues(fields []*Field) Set {
	validFieldValues := make(Set)
	for _, field := range fields {
		validFieldValues.Union(field.Set)
	}
	return validFieldValues
}
