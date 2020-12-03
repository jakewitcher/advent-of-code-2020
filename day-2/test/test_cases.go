package validate

import "day-2/internal/validate"

var parseMinMaxTestCases = []struct {
	min, max int
	input    string
	success  bool
}{
	{
		min:     1,
		max:     2,
		input:   "1-2",
		success: true,
	},
	{
		min:     10,
		max:     12,
		input:   "10-12",
		success: true,
	},
	{
		min:     -1,
		max:     -1,
		input:   "ab-cd",
		success: false,
	},
}

var parseCharacterTestCases = []struct {
	character rune
	input     string
	success   bool
}{
	{
		character: 'b',
		input:     "b:",
		success:   true,
	},
	{
		character: '0',
		input:     "2:",
		success:   false,
	},
}

var parsePasswordPolicyTestCases = []struct {
	password validate.Password
	policy   *validate.PasswordPolicy
	input    string
	success  bool
}{
	{
		password: validate.Password("cvjcc"),
		policy: &validate.PasswordPolicy{Character: 'c', Min: 1, Max: 3},
		input: "1-3 c: cvjcc",
		success: true,
	},
	{
		password: validate.Password("nkfgnr"),
		policy: &validate.PasswordPolicy{Character: 'r', Min: 10, Max: 13},
		input: "10-13 r: nkfgnr",
		success: true,
	},
	{
		password: validate.Password(""),
		policy: nil,
		input: "1-3 r: nkfgnr 45",
		success: false,
	},
	{
		password: validate.Password(""),
		policy: nil,
		input: "ab-cd r: nkfgnr",
		success: false,
	},
	{
		password: validate.Password(""),
		policy: nil,
		input: "1-3-7 r: nkfgnr",
		success: false,
	},
	{
		password: validate.Password(""),
		policy: nil,
		input: "1-3 45: nkfgnr",
		success: false,
	},
}

var identifyValidSledRentalPasswordsTestCases = []struct {
	expected int
	input []string
}{
	{
		expected: 1,
		input: []string{"1-3 a: abcdea"},
	},
	{
		expected: 0,
		input: []string{"3-5 a: abcdea"},
	},
	{
		expected: 2,
		input: []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"},
	},
}

var identifyValidTobogganCorpPasswordsTestCases = []struct {
	expected int
	input []string
}{
	{
		expected: 1,
		input: []string{"1-3 a: abcdea"},
	},
	{
		expected: 0,
		input: []string{"3-5 a: abcdea"},
	},
	{
		expected: 1,
		input: []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"},
	},
}
