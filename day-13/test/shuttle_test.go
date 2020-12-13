package shuttle

import (
	"day-13/internal/input"
	"day-13/internal/shuttle"
	"testing"
)

func TestFindNextAvailableBus(t *testing.T) {
	for _, test := range nextAvailableBusTestCases {
		actual, err := shuttle.FindNextAvailableBus(test.timestamp, test.input)
		if err != nil {
			t.Fatalf("error not expected, %s", err.Error())
		}

		if actual != test.expected {
			t.Fatalf("expected %d, actual %d", test.expected, actual)
		}
	}
}

func TestFindEarliestTimestamp(t *testing.T) {
	for _, test := range earliestTimestampTestCases {
		actual, err := shuttle.FindEarliestTimestamp(test.input)
		if err != nil {
			t.Fatalf("error not expected, %s", err.Error())
		}

		if actual != test.expected {
			t.Fatalf("expected %d, actual %d", test.expected, actual)
		}
	}
}

func BenchmarkFindNextAvailableBus(b *testing.B) {
	ts, n, err := input.Extract("../internal/input/input.txt")
	if err != nil {
		b.Fatalf("failed reading input from file, %s", err.Error())
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = shuttle.FindNextAvailableBus(ts, n)
	}
}

func BenchmarkFindEarliestTimestamp(b *testing.B) {
	_, n, err := input.Extract("../internal/input/input.txt")
	if err != nil {
		b.Fatalf("failed reading input from file, %s", err.Error())
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = shuttle.FindEarliestTimestamp(n)
	}
}
