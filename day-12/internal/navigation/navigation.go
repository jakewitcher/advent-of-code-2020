package navigation

import (
	"math"
	"strconv"
)

type Instruction struct {
	Action rune
	Units  int
}

type Point struct {
	X, Y int
}

type Ship struct {
	Facing   rune
	Waypoint *Point
	Point
	Interpreter
}

func (s *Ship) RunInstructions(instructions []*Instruction) {
	for _, instruction := range instructions {
		s.RunInstruction(s, instruction)
	}
}

func FindManhattanDistance(input []string, interpreter Interpreter) (int, error) {
	instructions := make([]*Instruction, len(input))
	for i, line := range input {
		instruction, err := ParseInstruction(line)
		if err != nil {
			return 0, err
		}

		instructions[i] = instruction
	}

	ship := Ship{
		Facing:      'E',
		Waypoint:    &Point{X: 10, Y: 1},
		Interpreter: interpreter,
	}

	ship.RunInstructions(instructions)

	distance := math.Abs(float64(ship.X)) + math.Abs(float64(ship.Y))
	return int(distance), nil
}

func ParseInstruction(line string) (*Instruction, error) {
	r := rune(line[0])
	units, err := strconv.Atoi(line[1:])
	if err != nil {
		return nil, err
	}

	return &Instruction{Action: r, Units: units}, nil
}
