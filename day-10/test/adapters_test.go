package adapters

import (
	"day-10/internal/adapters"
	"day-10/internal/input"
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

func BenchmarkCalculatePaths(b *testing.B) {
	n, err := input.Extract("../internal/input/input.txt")
	if err != nil {
		b.Fatalf("failed reading input from file, %s", err.Error())
	}

	for i := 0; i < b.N; i++ {
		_ = adapters.CalculatePaths(n)
	}
}