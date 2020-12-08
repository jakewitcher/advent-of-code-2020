package game

import (
	"day-8/internal/game"
	"testing"
)

func TestParseInstruction(t *testing.T) {
	for _, test := range parseInstructionTestCases {
		actual, err := game.ParseInstruction(test.input)
		if err != nil && test.success {
			t.Fatalf("error not expected: %s", err.Error())
		}

		if err == nil && actual != test.expected {
			t.Fatalf("expected %v, actual %v", test.expected, actual)
		}
	}
}

func TestFindFinalAccumulatorBeforeLoop(t *testing.T) {
	for _, test := range findFinalAccumulatorBeforeLoopTestCases {
		actual, err := game.FindFinalAccumulatorBeforeLoop(test.input)
		if err != nil && test.success {
			t.Fatalf("error not expected: %s", err.Error())
		}

		if err == nil && actual != test.expected {
			t.Fatalf("expected %v, actual %v", test.expected, actual)
		}
	}
}

func TestFindFinalAccumulatorOfFixedProgram(t *testing.T) {
	for _, test := range findFinalAccumulatorOfFixedProgramTestCases {
		actual, err := game.FindFinalAccumulatorOfFixedProgram(test.input)
		if err != nil && test.success {
			t.Fatalf("error not expected: %s", err.Error())
		}

		if err == nil && actual != test.expected {
			t.Fatalf("expected %v, actual %v", test.expected, actual)
		}
	}
}
