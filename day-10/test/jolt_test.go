package jolt

import (
	"day-10/internal/jolt"
	"testing"
)

func TestCalculateDifference(t *testing.T) {
	for _, test := range testCases {
		if actual := jolt.CalculateDifference(test.input); actual != test.expectedDifference {
			t.Fatalf("expected %d, actual %d", test.expectedDifference, actual)
		}
	}
}

func TestCalculatePaths(t *testing.T) {
	for _, test := range testCases {
		if actual := jolt.CalculatePaths(test.input); actual != test.Paths {
			t.Fatalf("expected %d, actual %d", test.Paths, actual)
		}
	}
}
