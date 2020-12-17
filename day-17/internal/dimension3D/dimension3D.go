package dimension3D

const ACTIVE = '#'

type Point struct {
	X, Y, Z int
}

type State struct {
	IsActive  bool
	Neighbors []Point
}

type Dimension map[Point]State

func (d Dimension) CountActive() int {
	var sum int
	for _, state := range d {
		if state.IsActive {
			sum++
		}
	}

	return sum
}

func FindActiveCubesAfterNCycles(input [][]rune, cycles int) (int, error) {
	dimension := initialize(input)

	for i := 0; i < cycles; i++ {
		dimension = Cycle(dimension)
	}

	return dimension.CountActive(), nil
}

func Cycle(dimension Dimension) Dimension {
	nextDimension := make(Dimension)
	for _, state := range dimension {
		for _, neighbor := range state.Neighbors {
			if _, ok := dimension[neighbor]; ok {
				continue
			}

			neighborsOfNeighbor := getNeighbors(neighbor)
			neighborState := State{
				IsActive:  false,
				Neighbors: neighborsOfNeighbor,
			}

			dimension[neighbor] = neighborState
		}
	}

	for point, state := range dimension {
		var activeNeighbors int

		for _, neighbor := range state.Neighbors {
			if neighborState, ok := dimension[neighbor]; ok && neighborState.IsActive {
				activeNeighbors++
			}
		}

		if state.IsActive {
			var isActive bool
			if activeNeighbors > 1 && activeNeighbors < 4 {
				isActive = true
			}

			nextDimension[point] = State{
				IsActive:  isActive,
				Neighbors: state.Neighbors,
			}

			continue
		}

		if !state.IsActive {
			var isActive bool
			if activeNeighbors == 3 {
				isActive = true
			}

			nextDimension[point] = State{
				IsActive:  isActive,
				Neighbors: state.Neighbors,
			}
			continue
		}

		nextDimension[point] = state
	}

	return nextDimension
}

func initialize(input [][]rune) Dimension {
	dimension := make(Dimension)

	for i, row := range input {
		for j, r := range row {
			var isActive bool
			if r == ACTIVE {
				isActive = true
			}

			point := Point{X: j, Y: i, Z: 0}
			neighbors := getNeighbors(point)
			state := State{
				IsActive:  isActive,
				Neighbors: neighbors,
			}
			dimension[point] = state
		}
	}

	return dimension
}

func getNeighbors(p Point) []Point {
	neighbors := make([]Point, 26)
	var index int
	for x := p.X - 1; x <= p.X+1; x++ {
		for y := p.Y - 1; y <= p.Y+1; y++ {
			for z := p.Z - 1; z <= p.Z+1; z++ {
				neighbor := Point{X: x, Y: y, Z: z}
				if neighbor == p {
					continue
				}

				neighbors[index] = neighbor
				index++
			}
		}
	}

	return neighbors
}
