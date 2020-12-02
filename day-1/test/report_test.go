package report

import (
	"day-1/internal/report"
	"testing"
)

func TestFindProductOfTwoEntriesWithSum(t *testing.T) {
	for _, test := range twoEntriesTestCases {
		actual, err := report.FindProductOfTwoEntriesWithSum(test.expenses, test.target)
		if err != nil && test.success {
			t.Fatal("test case should have produced a valid pair of entries")
		}

		if actual != test.expected {
			t.Fatalf("expected: %v, actual: %v", test.expected, actual)
		}
	}
}

func TestFindProductOfThreeEntriesWithSum(t *testing.T) {
	for _, test := range threeEntriesTestCases {
		actual, err := report.FindProductOfThreeEntriesWithSum(test.expenses, test.target)
		if err != nil && test.success {
			t.Fatal("test case should have produced a valid pair of entries")
		}

		if actual != test.expected {
			t.Fatalf("expected: %v, actual: %v", test.expected, actual)
		}
	}
}