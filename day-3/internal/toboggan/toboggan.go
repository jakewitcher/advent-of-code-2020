package toboggan

type Matrix [][]rune

type Slope struct {
	Right, Down int
}

type Location struct {
	x, y int
}

type Area struct {
	width, height int
}

type SleddingMap struct {
	Area
	Location
	Slope
}

func (s *SleddingMap) Move() {
	if s.x+s.Right > s.width {
		s.x = ((s.x + s.Right) % s.width) - 1
	} else {
		s.x += s.Right
	}

	s.y += s.Down
}

func (s *SleddingMap) IsOutOfBounds() bool {
	return s.y > s.height
}

func EvaluateSleddingPath(matrix Matrix, slope Slope) int {
	width := len(matrix[0]) - 1
	height := len(matrix) - 1
	area := Area{width: width, height: height}
	sled := SleddingMap{Area: area, Slope: slope}

	var trees int
	for !sled.IsOutOfBounds() {
		r := matrix[sled.y][sled.x]
		if r == '#' {
			trees++
		}

		sled.Move()
	}

	return trees
}
