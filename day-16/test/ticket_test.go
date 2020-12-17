package ticket

import (
	"day-16/internal/input"
	"day-16/internal/ticket"
	"testing"
)

func TestParseField(t *testing.T) {
	for _, test := range parseFieldTestCaes {
		actual, err := ticket.ParseField(test.input)
		if err != nil {
			t.Fatalf("no error expected, %s", err.Error())
		}

		if actual.Name != test.name {
			t.Fatalf("expected name %s, actual %s", test.name, actual.Name)
		}

		if len(actual.Set) != len(test.set) {
			t.Fatalf("expected set of length %d, actual %d", len(test.set), len(actual.Set))
		}

		for _, n := range test.set {
			if _, ok := actual.Set[n]; !ok {
				t.Fatalf("expected %d not found in set", n)
			}
		}
	}
}

func BenchmarkScanErrorRate(b *testing.B) {
	fields, _, nearbyTickets, err := input.Extract("../internal/input/input.txt")
	if err != nil {
		b.Fatalf("unexpected error, %s", err.Error())
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = ticket.ScanErrorRate(fields, nearbyTickets)
	}
}

func BenchmarkIdentifyDepartureFields(b *testing.B) {
	fields, yourTicket, nearbyTickets, err := input.Extract("../internal/input/input.txt")
	if err != nil {
		b.Fatalf("unexpected error, %s", err.Error())
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = ticket.IdentifyDepartureFields(fields, yourTicket, nearbyTickets)
	}
}
