package seating

import (
	"fmt"
)

const (
	OCCUPIED = '#'
	EMPTY    = 'L'
)

type Seating struct {
	Width, Height int
	Matrix        [][]rune
	Rules
}

func Normalize(matrix [][]rune, rules Rules) (int, error) {
	seating := initialize(matrix, rules)
	changes := seating.Height * seating.Width

	for changes > 0 {
		newSeating, newChanges, err := seating.ApplyRules()
		if err != nil {
			return changes, err
		}

		changes = newChanges
		seating = newSeating
	}

	return seating.countOccupied(), nil
}

func (s Seating) ApplyRules() (Seating, int, error) {
	newSeating := createEmpty(s.Rules, s.Height, s.Width)
	var changes int

	for i, row := range s.Matrix {
		for j, current := range row {
			point := Point{X: j, Y: i}

			if current == OCCUPIED && s.IsDeterrent(s, point) {
				err := newSeating.set(point, EMPTY)
				if err != nil {
					return newSeating, changes, err
				}
				changes++
				continue
			}

			if current == EMPTY && s.IsIncentive(s, point) {
				err := newSeating.set(point, OCCUPIED)
				if err != nil {
					return newSeating, changes, err
				}
				changes++
				continue
			}

			err := newSeating.set(point, current)
			if err != nil {
				return newSeating, changes, err
			}
		}
	}

	return newSeating, changes, nil
}

func (s Seating) countOccupied() int {
	var count int
	for _, row := range s.Matrix {
		for _, current := range row {
			if current == OCCUPIED {
				count++
			}
		}
	}

	return count
}

func (s Seating) isOccupied(point Point) bool {
	r, err := s.get(point)
	if err != nil {
		return false
	}

	return r == OCCUPIED
}

func (s Seating) isEmpty(point Point) bool {
	r, err := s.get(point)
	if err != nil {
		return false
	}

	return r == EMPTY
}

func (s Seating) get(point Point) (rune, error) {
	if !s.isWithinRange(point) {
		return ' ', fmt.Errorf("point %v out of range", point)
	}

	return s.Matrix[point.Y][point.X], nil
}

func (s Seating) set(point Point, r rune) error {
	if !s.isWithinRange(point) {
		return fmt.Errorf("point %v out of range", point)
	}

	s.Matrix[point.Y][point.X] = r

	return nil
}

func (s Seating) isWithinRange(point Point) bool {
	x, y := point.X, point.Y
	return x >= 0 && x < s.Width && y >= 0 && y < s.Height
}

func initialize(matrix [][]rune, rules Rules) Seating {
	seating := Seating{
		Height: len(matrix),
		Width:  len(matrix[0]),
		Matrix: matrix,
		Rules:  rules,
	}
	return seating
}

func createEmpty(rules Rules, height, width int) Seating {
	matrix := make([][]rune, height)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]rune, width)
	}

	return Seating{
		Height: height,
		Width:  width,
		Matrix: matrix,
		Rules:  rules,
	}
}
