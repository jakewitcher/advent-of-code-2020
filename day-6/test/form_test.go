package form

import (
	"day-6/internal/form"
	"day-6/internal/input"
	"testing"
)

func TestEvaluateTotalYesResponsesByGroup(t *testing.T) {
	for _, test := range testCases {
		if total := form.EvaluateTotalYesResponsesByGroup(test.input); total != test.total {
			t.Fatalf("total %d, total %d", test.total, total)
		}
	}
}

func TestEvaluateConsensusYesResponsesByGroup(t *testing.T) {
	for _, test := range testCases {
		if consensus := form.EvaluateConsensusYesResponsesByGroup(test.input); consensus != test.consensus {
			t.Fatalf("consensus %d, consensus %d", test.consensus, consensus)
		}
	}
}

func BenchmarkEvaluateTotalYesResponsesByGroup(b *testing.B) {
	n, err := input.Extract("../internal/input/input.txt")
	if err != nil {
		b.Fatalf("failed reading input from file, %s", err.Error())
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		group := i % (len(n) - 1)
		_ = form.EvaluateTotalYesResponsesByGroup(n[group])
	}
}

func BenchmarkEvaluateConsensusYesResponsesByGroup(b *testing.B) {
	n, err := input.Extract("../internal/input/input.txt")
	if err != nil {
		b.Fatalf("failed reading input from file, %s", err.Error())
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		group := i % (len(n) - 1)
		_ = form.EvaluateConsensusYesResponsesByGroup(n[group])
	}
}
