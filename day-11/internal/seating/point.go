package seating

const (
	N = iota
	NE
	E
	SE
	S
	SW
	W
	NW
)

type Point struct {
	X, Y int
}

type Direction int

func (p Point) Move(direction Direction) Point {
	var x, y int
	switch direction {
	case N:
		x, y = p.X, p.Y-1
	case NE:
		x, y = p.X+1, p.Y-1
	case E:
		x, y = p.X+1, p.Y
	case SE:
		x, y = p.X+1, p.Y+1
	case S:
		x, y = p.X, p.Y+1
	case SW:
		x, y = p.X-1, p.Y+1
	case W:
		x, y = p.X-1, p.Y
	case NW:
		x, y = p.X-1, p.Y-1
	}

	return Point{x, y}
}
