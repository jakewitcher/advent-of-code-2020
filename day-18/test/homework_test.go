package homework

import (
	"day-18/internal/homework"
	"day-18/internal/input"
	"testing"
)

func TestParseExpressionEqualPrecedence(t *testing.T) {
	for _, test := range parseExpressionEqualPrecedenceTestCases {
		actual, err := homework.ParseExpression(test.input, homework.EqualPrecedence{})
		if err != nil {
			t.Fatalf("error not expected, %s", err.Error())
		}

		if actual.Evaluate() != test.expected {
			t.Fatalf("expected %d, actual %d", test.expected, actual.Evaluate())
		}
	}
}

func TestParseExpressionAdditionHigherPrecedence(t *testing.T) {
	for _, test := range parseExpressionAdditionHigherPrecedenceTestCases {
		actual, err := homework.ParseExpression(test.input, homework.AdditionHigherPrecedence{})
		if err != nil {
			t.Fatalf("error not expected, %s", err.Error())
		}

		if actual.Evaluate() != test.expected {
			t.Fatalf("expected %d, actual %d", test.expected, actual.Evaluate())
		}
	}
}

func BenchmarkSumEvaluatedExpressionsEqualPrecedence(b *testing.B) {
	n, err := input.Extract("../internal/input/input.txt")
	if err != nil {
		b.Fatalf("error not expected, %s", err.Error())
	}

	for i := 0; i < b.N; i++ {
		_, _ = homework.SumEvaluatedExpressions(n, homework.EqualPrecedence{})
	}
}

func BenchmarkSumEvaluatedExpressionsAdditionHigherPrecedence(b *testing.B) {
	n, err := input.Extract("../internal/input/input.txt")
	if err != nil {
		b.Fatalf("error not expected, %s", err.Error())
	}

	for i := 0; i < b.N; i++ {
		_, _ = homework.SumEvaluatedExpressions(n, homework.AdditionHigherPrecedence{})
	}
}
