package game

import "day-8/internal/game"

var parseInstructionTestCases = []struct {
	input    string
	expected game.Instruction
	success  bool
}{
	{
		input: "nop +456",
		expected: game.Instruction{
			Operation: "nop",
			Argument:  456,
		},
		success: true,
	},
	{
		input: "acc +9",
		expected: game.Instruction{
			Operation: "acc",
			Argument:  9,
		},
		success: true,
	},
	{
		input: "jmp -153",
		expected: game.Instruction{
			Operation: "jmp",
			Argument:  -153,
		},
		success: true,
	},
}

var findFinalAccumulatorBeforeLoopTestCases = []struct {
	input    []string
	expected int
	success  bool
}{
	{
		input: []string{
			"nop +0",
			"acc +1",
			"jmp +4",
			"acc +3",
			"jmp -3",
			"acc -99",
			"acc +1",
			"jmp -4",
			"acc +6",
		},
		expected: 5,
		success:  true,
	},
}

var findFinalAccumulatorOfFixedProgramTestCases = []struct {
	input    []string
	expected int
	success  bool
}{
	{
		input: []string{
			"nop +0",
			"acc +1",
			"jmp +4",
			"acc +3",
			"jmp -3",
			"acc -99",
			"acc +1",
			"jmp -4",
			"acc +6",
		},
		expected: 8,
		success:  true,
	},
}
