package validate

import (
	"day-2/internal/input"
	"testing"
)
import "day-2/internal/validate"

func TestParsePolicyMinMax(t *testing.T) {
	for _, test := range parseMinMaxTestCases {
		min, max, err := validate.ParsePolicyMinMax(test.input)
		if err != nil && test.success {
			t.Errorf("error not expected: %s", err.Error())
		}

		if test.success && (min != test.min || max != test.max) {
			t.Errorf("expected (%d, %d) actual (%d, %d)", test.min, test.max, min, max)
		}
	}
}

func TestParsePolicyCharacter(t *testing.T) {
	for _, test := range parseCharacterTestCases {
		character, err := validate.ParsePolicyCharacter(test.input)
		if err != nil && test.success {
			t.Errorf("error not expected: %s", err.Error())
		}

		if test.success && character != test.character {
			t.Errorf("expected %c, actual %c", test.character, character)
		}
	}
}

func TestParsePasswordPolicy(t *testing.T) {
	for _, test := range parsePasswordPolicyTestCases {
		password, policy, err := validate.ParsePasswordPolicy(test.input)
		if err != nil && test.success {
			t.Errorf("error not expected: %s", err.Error())
		}

		if err == nil && password != test.password {
			t.Errorf("expected %s, actual %s", test.password, password)
		}

		if err == nil && *policy != *test.policy {
			t.Errorf("expected %#v, actual %#v", test.policy, policy)
		}
	}
}

func TestIdentifyValidPasswordsSledRentalPolicy(t *testing.T) {
	for _, test := range identifyValidSledRentalPasswordsTestCases {
		actual, err := validate.IdentifyValidPasswords(test.input, validate.ApplySledRentalPolicy)
		if err != nil {
			t.Errorf("error not expected: %s", err.Error())
			continue
		}

		if actual != test.expected {
			t.Errorf("expected %d, actual %d", test.expected, actual)
		}
	}
}

func TestIdentifyValidPasswordsTobogganCorpPolicy(t *testing.T) {
	for _, test := range identifyValidTobogganCorpPasswordsTestCases {
		actual, err := validate.IdentifyValidPasswords(test.input, validate.ApplyTobogganCorpPolicy)
		if err != nil {
			t.Errorf("error not expected: %s", err.Error())
			continue
		}

		if actual != test.expected {
			t.Errorf("expected %d, actual %d", test.expected, actual)
		}
	}
}

func BenchmarkIdentifyValidPasswordsSledRentalPolicy(b *testing.B) {
	n, err := input.Extract("../internal/input/input.txt")
	if err != nil {
		b.Fatalf("failed reading input from file, %s", err.Error())
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = validate.IdentifyValidPasswords(n, validate.ApplySledRentalPolicy)
	}
}

func BenchmarkIdentifyValidPasswordsTobogganCorpPolicy(b *testing.B) {
	n, err := input.Extract("../internal/input/input.txt")
	if err != nil {
		b.Fatalf("failed reading input from file, %s", err.Error())
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = validate.IdentifyValidPasswords(n, validate.ApplyTobogganCorpPolicy)
	}
}
