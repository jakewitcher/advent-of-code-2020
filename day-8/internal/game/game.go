package game

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Instruction struct {
	Operation string
	Argument  int
}

func (i *Instruction) Flip() {
	if i.Operation == "jmp" {
		i.Operation = "nop"
	} else if i.Operation == "nop" {
		i.Operation = "jmp"
	}
}

func FindFinalAccumulatorOfFixedProgram(input []string) (int, error) {
	instructions, err := ParseInstructions(input)
	if err != nil {
		return 0, err
	}

	operations := make(map[int]interface{})
	for i, instruction := range instructions {
		if instruction.Operation == "jmp" || instruction.Operation == "nop" {
			operations[i] = struct{}{}
		}
	}

	for i, _ := range operations {
		instructions[i].Flip()
		accumulator, finalPosition, err := RunBootCode(instructions)
		if err != nil {
			return 0, err
		}

		if finalPosition == len(instructions) {
			return accumulator, nil
		}

		instructions[i].Flip()
	}

	return 0, errors.New("no fix found")
}

func FindFinalAccumulatorBeforeLoop(input []string) (int, error) {
	instructions, err := ParseInstructions(input)
	if err != nil {
		return 0, err
	}

	accumulator, _, err := RunBootCode(instructions)
	if err != nil {
		return 0, err
	}

	return accumulator, nil
}

func RunBootCode(instructions []Instruction) (int, int, error) {
	var accumulator, i int
	var err error

	seen := make(map[int]interface{})

	for {
		if i > len(instructions)-1 {
			break
		}

		if _, ok := seen[i]; ok {
			break
		}

		seen[i] = struct{}{}
		instruction := instructions[i]

		accumulator, i, err = RunInstruction(instruction, accumulator, i)
		if err != nil {
			return 0, 0, err
		}
	}

	return accumulator, i, nil
}

func RunInstruction(instruction Instruction, accumulator, i int) (int, int, error) {
	switch instruction.Operation {
	case "acc":
		return accumulator + instruction.Argument, i + 1, nil
	case "jmp":
		return accumulator, i + instruction.Argument, nil
	case "nop":
		return accumulator, i + 1, nil
	default:
		return accumulator, i, fmt.Errorf("not a valid operation: %s", instruction.Operation)
	}
}

func ParseInstructions(inputs []string) ([]Instruction, error) {
	instructions := make([]Instruction, len(inputs))

	for i, input := range inputs {
		instruction, err := ParseInstruction(input)
		if err != nil {
			return nil, err
		}

		instructions[i] = instruction
	}

	return instructions, nil
}

func ParseInstruction(input string) (Instruction, error) {
	components := strings.Split(input, " ")
	op := components[0]
	arg, err := strconv.Atoi(components[1][1:])
	if err != nil {
		return Instruction{}, err
	}

	if components[1][0] == '-' {
		arg = -arg
	}

	return Instruction{Operation: op, Argument: arg}, nil
}
