package dimension

import (
	"day-17/internal/dimension3D"
	"day-17/internal/dimension4D"
	"day-17/internal/input"
	"log"
	"testing"
)

func BenchmarkFindActiveCubesAfterNCycles3D(b *testing.B) {
	cubes, err := input.Extract("../internal/input/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = dimension3D.FindActiveCubesAfterNCycles(cubes, 6)
	}
}

func BenchmarkFindActiveCubesAfterNCycles4D(b *testing.B) {
	cubes, err := input.Extract("../internal/input/input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = dimension4D.FindActiveCubesAfterNCycles(cubes, 6)
	}
}
