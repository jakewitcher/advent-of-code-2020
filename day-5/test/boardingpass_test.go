package boardingpass

import (
	"day-5/internal/boardingpass"
	"testing"
)

func TestParseRow(t *testing.T) {
	for _, test := range parseRowTestCases {
		if actual := boardingpass.ParseRow(test.input); actual != test.expected {
			t.Fatalf("expected %d, actual %d", test.expected, actual)
		}
	}
}

func TestParseColumn(t *testing.T) {
	for _, test := range parseColumnTestCases {
		if actual := boardingpass.ParseColumn(test.input); actual != test.expected {
			t.Fatalf("expected %d, actual %d", test.expected, actual)
		}
	}
}

func TestParseBoardingPass(t *testing.T) {
	for _, test := range parseBoardingPassTestCases {
		if actual := boardingpass.Parse(test.input); actual.SeatId != test.expected {
			t.Fatalf("expected %d, actual %d", test.expected, actual.SeatId)
		}
	}
}
