package luggage

import (
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

func TestSumPossibleOuterBags(t *testing.T) {
	for _, test := range sumOuterBagsTestCases {
		actual, err := luggage.SumPossibleOuterBags(test.input, "shiny gold")
		if err != nil && test.success {
			t.Fatalf("error not expected: %s", err.Error())
		}

		if err == nil && actual != test.expected {
			t.Fatalf("expected %d, actual %d", test.expected, actual)
		}
	}
}

func TestSumRequiredInnerBags(t *testing.T) {
	for _, test := range sumInnerBagsTestCases {
		actual, err := luggage.SumRequiredInnerBags(test.input, "shiny gold")
		if err != nil && test.success {
			t.Fatalf("error not expected: %s", err.Error())
		}

		if err == nil && actual != test.expected {
			t.Fatalf("expected %d, actual %d", test.expected, actual)
		}
	}
}
