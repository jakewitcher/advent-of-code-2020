package ticket

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ParseTickets(input []string) ([]Ticket, error) {
	tickets := make([]Ticket, len(input))
	for i, ticketStr := range input {
		ticket, err := ParseTicket(ticketStr)
		if err != nil {
			return nil, err
		}

		tickets[i] = ticket
	}

	return tickets, nil
}

func ParseTicket(input string) (Ticket, error) {
	fieldValuesStr := strings.Split(input, ",")
	fieldValues := make([]int, len(fieldValuesStr))
	for i, fieldValueStr := range fieldValuesStr {
		fieldValue, err := strconv.Atoi(fieldValueStr)
		if err != nil {
			return nil, err
		}

		fieldValues[i] = fieldValue
	}

	return fieldValues, nil
}

func ParseFields(input []string) ([]*Field, error) {
	fields := make([]*Field, len(input))

	for i, fieldStr := range input {
		field, err := ParseField(fieldStr)
		if err != nil {
			return nil, err
		}

		fields[i] = &field
	}

	return fields, nil
}

func ParseField(input string) (Field, error) {
	set, err := parseFieldSet(input)
	if err != nil {
		return Field{}, err
	}

	name, err := parseFieldName(input)
	if err != nil {
		return Field{}, err
	}

	return Field{Name: name, Set: set}, nil
}

func parseFieldName(input string) (string, error) {
	regex, err := regexp.Compile("^[\\w\\s]*")
	if err != nil {
		return "", err
	}

	name := regex.Find([]byte(input))
	return string(name), nil
}

func parseFieldSet(input string) (Set, error) {
	regex, err := regexp.Compile("\\d+")
	if err != nil {
		return nil, err
	}

	rangeValuesStr := regex.FindAll([]byte(input), 4)
	if len(rangeValuesStr) != 4 {
		return nil, fmt.Errorf("invalid field format: %s", input)
	}

	rangeValues := make([]int, len(rangeValuesStr))
	for i, valueStr := range rangeValuesStr {
		value, err := strconv.Atoi(string(valueStr))
		if err != nil {
			return nil, err
		}

		rangeValues[i] = value
	}

	set := make(Set)
	for i := rangeValues[0]; i <= rangeValues[1]; i++ {
		set[i] = struct{}{}
	}

	for i := rangeValues[2]; i <= rangeValues[3]; i++ {
		set[i] = struct{}{}
	}
	return set, nil
}
