package toboggan

import (
	"day-3/internal/input"
	"day-3/internal/toboggan"
	"testing"
)

func BenchmarkEvaluateSleddingPath(b *testing.B) {
	n, err := input.Extract("../internal/input/input.txt")
	if err != nil {
		b.Fatalf("failed reading input from file, %s", err.Error())
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		slope := toboggan.Slope{Right: 3, Down: 1}
		_ = toboggan.EvaluateSleddingPath(n, slope)
	}
}
