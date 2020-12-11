package report

import (
	"day-1/internal/input"
	"day-1/internal/report"
	"testing"
)

func TestFindProductOfTwoEntriesWithTargetSum(t *testing.T) {
	for _, test := range twoEntriesTestCases {
		actual, err := report.FindProductOfTwoEntriesWithTargetSum(test.expenses, test.target)
		if err != nil && test.success {
			t.Fatal("test case should have produced a valid pair of entries")
		}

		if actual != test.expected {
			t.Fatalf("expected: %d, actual: %d", test.expected, actual)
		}
	}
}

func TestFindProductOfThreeEntriesWithSum(t *testing.T) {
	for _, test := range threeEntriesTestCases {
		actual, err := report.FindProductOfThreeEntriesWithTargetSum(test.expenses, test.target)
		if err != nil && test.success {
			t.Fatal("test case should have produced a valid pair of entries")
		}

		if actual != test.expected {
			t.Fatalf("expected: %d, actual: %d", test.expected, actual)
		}
	}
}

func BenchmarkFindProductOfTwoEntriesWithTargetSum(b *testing.B) {
	n, err := input.Extract("../internal/input/input.txt")
	if err != nil {
		b.Fatalf("failed reading input from file, %s", err.Error())
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = report.FindProductOfTwoEntriesWithTargetSum(n, 2020)
	}
}

func BenchmarkFindProductOfThreeEntriesWithTargetSum(b *testing.B) {
	n, err := input.Extract("../internal/input/input.txt")
	if err != nil {
		b.Fatalf("failed reading input from file, %s", err.Error())
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = report.FindProductOfThreeEntriesWithTargetSum(n, 2020)
	}
}

