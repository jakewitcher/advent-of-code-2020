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
	rng := 128

	for _, r := range input {
		rng /= 2
		if r == 'F' {
			max -= rng
		} else if r == 'B' {
			min += rng
		}
	}

	return max
}

func ParseColumn(input string) int {
	min := 0
	max := 7
	rng := 8

	for _, r := range input {
		rng /= 2
		if r == 'L' {
			max -= rng
		} else if r == 'R' {
			min += rng
		}
	}

	return max
}

