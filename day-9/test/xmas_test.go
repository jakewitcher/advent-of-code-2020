package xmas

import (
	"day-9/internal/xmas"
	"testing"
)

func TestFirstWeakness(t *testing.T) {
	for _, test := range testCases {
		actual, err := xmas.FindFirstWeakness(test.input, test.preamble)
		if err != nil && test.success {
			t.Fatalf("no error expected: %s", err.Error())
		}

		if err == nil && actual != test.firstExpected {
			t.Fatalf("expected %d, actual %d", test.firstExpected, actual)
		}
	}
}

func TestSecondWeakness(t *testing.T) {
	for _, test := range testCases {
		actual, err := xmas.FindSecondWeakness(test.input, test.preamble)
		if err != nil && test.success {
			t.Fatalf("no error expected: %s", err.Error())
		}

		if err == nil && actual != test.secondExpected {
			t.Fatalf("expected %d, actual %d", test.secondExpected, actual)
		}
	}
}
