package seating

type Rules interface {
	IsIncentive(Seating, Point) bool
	IsDeterrent(Seating, Point) bool
}

type AdjacentRules struct{}

func (AdjacentRules) IsIncentive(seating Seating, point Point) bool {
	return countAdjacentOccupied(seating, point) == 0
}

func (AdjacentRules) IsDeterrent(seating Seating, point Point) bool {
	return countAdjacentOccupied(seating, point) >= 4
}

func countAdjacentOccupied(seating Seating, point Point) int {
	var count int
	directions := [8]Direction{N, NE, E, SE, S, SW, W, NW}
	for _, direction := range directions {
		adj := point.Move(direction)
		if seating.isOccupied(adj) {
			count++
		}
	}

	return count
}

type VisibleRules struct{}

func (VisibleRules) IsIncentive(seating Seating, point Point) bool {
	return countVisibleOccupied(seating, point) == 0
}

func (VisibleRules) IsDeterrent(seating Seating, point Point) bool {
	return countVisibleOccupied(seating, point) >= 5
}

func countVisibleOccupied(seating Seating, point Point) int {
	directions := [8]Direction{N, NE, E, SE, S, SW, W, NW}
	var count int
	for _, direction := range directions {
		if isOccupiedVisible(seating, point, direction) {
			count++
		}
	}

	return count
}

func isOccupiedVisible(seating Seating, point Point, direction Direction) bool {
	for seating.isWithinRange(point) {
		point = point.Move(direction)
		if seating.isOccupied(point) {
			return true
		}

		if seating.isEmpty(point) {
			return false
		}
	}

	return false
}
