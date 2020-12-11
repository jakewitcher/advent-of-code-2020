package xmas

import (
	"day-9/internal/input"
	"day-9/internal/xmas"
	"testing"
)

func TestFindFirstWeakness(t *testing.T) {
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

func TestFindSecondWeakness(t *testing.T) {
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

func BenchmarkFindFirstWeakness(b *testing.B) {
	n, err := input.Extract("../internal/input/input.txt")
	if err != nil {
		b.Fatalf("failed reading input from file, %s", err.Error())
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = xmas.FindFirstWeakness(n, 25)
	}
}

func BenchmarkFindSecondWeakness(b *testing.B) {
	n, err := input.Extract("../internal/input/input.txt")
	if err != nil {
		b.Fatalf("failed reading input from file, %s", err.Error())
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = xmas.FindSecondWeakness(n, 25)
	}
}
