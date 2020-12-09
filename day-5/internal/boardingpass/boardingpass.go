package boardingpass

type BoardingPass struct {
	Row, Column int
	SeatId      int
}

func Parse(input string) *BoardingPass {
	row := ParseRow(input[:7])
	column := ParseColumn(input[7:])

	return &BoardingPass{Row: row, Column: column, SeatId: row*8 + column}
}

func ParseRow(input string) int {
	min := 0
	max := 127
	width := 128

	for _, r := range input {
		width /= 2
		if r == 'F' {
			max -= width
		} else if r == 'B' {
			min += width
		}
	}

	return max
}

func ParseColumn(input string) int {
	min := 0
	max := 7
	height := 8

	for _, r := range input {
		height /= 2
		if r == 'L' {
			max -= height
		} else if r == 'R' {
			min += height
		}
	}

	return max
}

