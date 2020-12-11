package luggage

import (
	"day-7/internal/input"
	"day-7/internal/luggage"
	"testing"
)

func TestParseRule(t *testing.T) {
	for _, test := range parseRuleTestCases {
		actual, err := luggage.ParseRule(test.input)
		if err != nil && test.success {
			t.Fatalf("error not expected: %s", err.Error())
		}

		if actual == nil || err != nil {
			continue
		}

		if actual.BagName != test.expected.BagName {
			t.Fatalf("expected %s, actual %s", test.expected.BagName, actual.BagName)
		}

		if len(actual.ContainedBags) != len(test.expected.ContainedBags) {
			t.Fatalf("expected %d contained bags, actual %d", len(test.expected.ContainedBags), len(actual.ContainedBags))
		}

		for i, bag := range test.expected.ContainedBags {
			if *bag != *actual.ContainedBags[i] {
				t.Fatalf("expected %v, actual %v", bag, actual.ContainedBags[i])
			}
		}
	}
}

func TestSumOuterBags(t *testing.T) {
	for _, test := range sumOuterBagsTestCases {
		actual, err := luggage.SumOuterBags(test.input, "shiny gold")
		if err != nil && test.success {
			t.Fatalf("error not expected: %s", err.Error())
		}

		if err == nil && actual != test.expected {
			t.Fatalf("expected %d, actual %d", test.expected, actual)
		}
	}
}

func TestSumInnerBags(t *testing.T) {
	for _, test := range sumInnerBagsTestCases {
		actual, err := luggage.SumInnerBags(test.input, "shiny gold")
		if err != nil && test.success {
			t.Fatalf("error not expected: %s", err.Error())
		}

		if err == nil && actual != test.expected {
			t.Fatalf("expected %d, actual %d", test.expected, actual)
		}
	}
}

func BenchmarkSumOuterBags(b *testing.B) {
	n, err := input.Extract("../internal/input/input.txt")
	if err != nil {
		b.Fatalf("failed reading input from file, %s", err.Error())
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = luggage.SumOuterBags(n, "shiny gold")
	}
}

func BenchmarkSumInnerBags(b *testing.B) {
	n, err := input.Extract("../internal/input/input.txt")
	if err != nil {
		b.Fatalf("failed reading input from file, %s", err.Error())
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = luggage.SumInnerBags(n, "shiny gold")
	}
}