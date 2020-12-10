package adapters

import (
	"day-10/internal/adapters"
	"testing"
)

func TestCalculateDifference(t *testing.T) {
	for _, test := range testCases {
		if actual := adapters.CalculateDifference(test.input); actual != test.expectedDifference {
			t.Fatalf("expected %d, actual %d", test.expectedDifference, actual)
		}
	}
}

func TestCalculatePaths(t *testing.T) {
	for _, test := range testCases {
		if actual := adapters.CalculatePaths(test.input); actual != test.Paths {
			t.Fatalf("expected %d, actual %d", test.Paths, actual)
		}
	}
}
