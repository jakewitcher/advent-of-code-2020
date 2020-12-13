package shuttle

import (
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

func TestFindMultipleOfXGreaterThanZWhereXPlusDiffEqualsMultipleOfY(t *testing.T) {
	for _, test := range findMultipleOfXTestCases {
		actual := shuttle.FindMultipleOfXGreaterThanZWhereXPlusOffsetEqualsMultipleOfY(test.x, test.y, test.z, test.offset)
		if actual != test.expected {
			t.Fatalf("expected %d, actual %d", test.expected, actual)
		}
	}
}

func BenchmarkFindEarliestTimestamp(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, test := range earliestTimestampTestCases {
			_, _ = shuttle.FindEarliestTimestamp(test.input)
		}
	}
}