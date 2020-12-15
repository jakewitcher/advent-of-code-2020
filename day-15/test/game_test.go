package game

import (
	"day-15/internal/game"
	"day-15/internal/input"
	"testing"
)

func TestFindNthNumberSpoken(t *testing.T) {
	for _, test := range testCases {
		if actual := game.FindNthNumberSpoken(test.input, test.target); actual != test.expected {
			t.Fatalf("expected %d, actual %d", test.expected, actual)
		}
	}
}

func BenchmarkFindNthNumberSpoken2020(b *testing.B) {
	n := input.Extract()

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = game.FindNthNumberSpoken(n, 2020)
	}
}

func BenchmarkFindNthNumberSpoken30Mil(b *testing.B) {
	n := input.Extract()

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = game.FindNthNumberSpoken(n, 30_000_000)
	}
}