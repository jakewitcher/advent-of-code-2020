package form

import (
	"day-6/internal/form"
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
