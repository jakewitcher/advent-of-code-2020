package ticket

import "sort"

type Accumulator map[string][]int

type FieldPositionIdentifier struct {
	Accumulator
	Fields []*Field
}

func NewPositionIdentifier(fields []*Field) FieldPositionIdentifier {
	identifier := FieldPositionIdentifier{
		Accumulator: make(Accumulator),
		Fields:      fields,
	}

	return identifier
}

func (t FieldPositionIdentifier) accumulateValidPositionsByField(tickets []Ticket) {
	validFieldValues := t.getValidFieldValues()

	for _, field := range t.Fields {
		t.Accumulator[field.Name] = make([]int, len(tickets[0]))
	}

	for _, ticket := range tickets {
		isValid := true
		for _, value := range ticket {
			if !validFieldValues.Contains(value) {
				isValid = false
				break
			}
		}

		if !isValid {
			continue
		}

		for _, field := range t.Fields {
			for i, value := range ticket {
				if !field.Contains(value) {
					t.Accumulator[field.Name][i]++
				}
			}
		}
	}
}

func (t FieldPositionIdentifier) getValidFieldValues() Set {
	validFieldValues := make(Set)
	for _, field := range t.Fields {
		validFieldValues.Union(field.Set)
	}
	return validFieldValues
}

func (t FieldPositionIdentifier) sortFieldsByNumberOfValidPositions() {
	for _, field := range t.Fields {
		var sum int
		for i, val := range t.Accumulator[field.Name] {
			if val == 0 {
				sum++
				field.ValidPositions = append(field.ValidPositions, i)
			}
		}

		field.TotalValid = sum
	}

	sort.Slice(t.Fields, func(a, b int) bool {
		return t.Fields[a].TotalValid < t.Fields[b].TotalValid
	})
}

func (t FieldPositionIdentifier) findActualPositionByField() map[string]int {
	fieldPositions := make(map[string]int)
	occupiedPositions := make(map[int]interface{})

	for _, field := range t.Fields {
		for _, validPosition := range field.ValidPositions {
			if _, ok := occupiedPositions[validPosition]; !ok {
				occupiedPositions[validPosition] = struct{}{}
				fieldPositions[field.Name] = validPosition
				break
			}
		}
	}

	return fieldPositions
}
