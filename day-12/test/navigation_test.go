package navigation

import (
	"day-12/internal/input"
	"day-12/internal/navigation"
	"testing"
)

func TestFindManhattanDistanceNaiveInterpreter(t *testing.T) {
	for _, test := range testCases {
		distance, err := navigation.FindManhattanDistance(test.input, navigation.NaiveInterpreter{})
		if err != nil {
			t.Fatalf("error not expected, %s", err.Error())
		}

		if distance != test.naiveDistance {
			t.Fatalf("expected %d, actual %d", test.naiveDistance, distance)
		}
	}
}

func TestFindManhattanDistanceWaypointInterpreter(t *testing.T) {
	for _, test := range testCases {
		distance, err := navigation.FindManhattanDistance(test.input, navigation.WaypointInterpreter{})
		if err != nil {
			t.Fatalf("error not expected, %s", err.Error())
		}

		if distance != test.waypointDistance {
			t.Fatalf("expected %d, actual %d", test.waypointDistance, distance)
		}
	}
}

func BenchmarkFindManhattanDistanceNaiveInterpreter(b *testing.B) {
	n, err := input.Extract("../internal/input/input.txt")
	if err != nil {
		b.Fatalf("failed reading input from file, %s", err.Error())
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = navigation.FindManhattanDistance(n, navigation.NaiveInterpreter{})
	}
}

func BenchmarkFindManhattanDistanceWaypointInterpreter(b *testing.B) {
	n, err := input.Extract("../internal/input/input.txt")
	if err != nil {
		b.Fatalf("failed reading input from file, %s", err.Error())
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = navigation.FindManhattanDistance(n, navigation.WaypointInterpreter{})
	}
}
