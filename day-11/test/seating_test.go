package seating

import (
	"day-11/internal/input"
	"day-11/internal/seating"
	"testing"
)

func TestApplyAdjacentRules(t *testing.T) {
	for _, test := range adjacentRulesTestCases {
		s := setupTestSeating(test.input, seating.AdjacentRules{})

		actual, changed, err := s.ApplyRules()
		if err != nil {
			t.Fatalf("error not expected, %s", err.Error())
		}

		if len(actual.Matrix) != len(test.expected) || len(actual.Matrix[0]) != len(test.expected[0]) {
			t.Fatalf("size of matrix is incorrect")
		}

		for i, row := range actual.Matrix {
			if string(row) != test.expected[i] {
				t.Fatalf("for row %d, expected %s, actual %s", i, test.expected[i], string(row))
			}
		}

		if changed != test.changed {
			t.Fatalf("expected %d changed, actual %d changed", test.changed, changed)
		}
	}
}

func TestApplyVisibleRules(t *testing.T) {
	for _, test := range visibleRulesTestCases {
		s := setupTestSeating(test.input, seating.VisibleRules{})

		actual, changed, err := s.ApplyRules()
		if err != nil {
			t.Fatalf("error not expected, %s", err.Error())
		}

		if len(actual.Matrix) != len(test.expected) || len(actual.Matrix[0]) != len(test.expected[0]) {
			t.Fatalf("size of matrix is incorrect")
		}

		for i, row := range actual.Matrix {
			if string(row) != test.expected[i] {
				t.Fatalf("for row %d, expected %s, actual %s", i, test.expected[i], string(row))
			}
		}

		if changed != test.changed {
			t.Fatalf("expected %d changed, actual %d changed", test.changed, changed)
		}
	}
}

func BenchmarkApplyAdjacentRules(b *testing.B) {
	n, err := input.Extract("../internal/input/input.txt")
	if err != nil {
		b.Fatalf("failed reading input from file, %s", err.Error())
	}

	s := setupSeating(n, seating.AdjacentRules{})

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = s.ApplyRules()
	}
}

func BenchmarkApplyVisibleRules(b *testing.B) {
	n, err := input.Extract("../internal/input/input.txt")
	if err != nil {
		b.Fatalf("failed reading input from file, %s", err.Error())
	}

	s := setupSeating(n, seating.VisibleRules{})

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = s.ApplyRules()
	}
}

func setupTestSeating(input []string, rules seating.Rules) seating.Seating {
	matrix := make([][]rune, len(input))
	for i, row := range input {
		matrix[i] = []rune(row)
	}

	return setupSeating(matrix, rules)
}

func setupSeating(matrix [][]rune, rules seating.Rules) seating.Seating {
	s := seating.Seating{
		Height: len(matrix),
		Width:  len(matrix[0]),
		Matrix: matrix,
		Rules:  rules,
	}

	return s
}
