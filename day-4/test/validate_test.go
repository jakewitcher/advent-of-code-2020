package validate

import (
	"day-4/internal/validate"
	"testing"
)

func TestParsePassportField(t *testing.T) {
	for _, test := range parsePassportFieldTestCases {
		actual, err := validate.ParsePassportField(test.input)
		if err != nil && test.success {
			t.Fatalf("error not expected, %s", err.Error())
		}

		if err == nil && *actual != test.expected {
			t.Fatalf("expected %#v, actual %#v", test.expected, *actual)
		}
	}
}

func TestParsePassport(t *testing.T) {
	for _, test := range parsePassportTestCases {
		actual, err := validate.ParsePassport(test.input)
		if err != nil && test.success {
			t.Fatalf("error not expected, %s", err.Error())
		}

		if err == nil && *actual != test.expected {
			t.Fatalf("expected %#v, actual %#v", test.expected, *actual)
		}
	}
}

func TestPassportIsValid(t *testing.T) {
	for _, test := range parsePassportIsValid {
		if isValid := test.passport.IsValid(); isValid != test.expected {
			t.Fatalf("expected %t, actual %t", test.expected, !test.expected)
		}
	}
}
