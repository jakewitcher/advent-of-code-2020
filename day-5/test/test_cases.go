package boardingpass

var parseRowTestCases = []struct {
	input    string
	expected int
}{
	{
		input:    "BFFFBBF",
		expected: 70,
	},
	{
		input:    "FFFBBBF",
		expected: 14,
	},
	{
		input:    "BBFFBBF",
		expected: 102,
	},
}

var parseColumnTestCases = []struct {
	input    string
	expected int
}{
	{
		input:    "RRR",
		expected: 7,
	},
	{
		input:    "LLL",
		expected: 0,
	},
	{
		input:    "RLL",
		expected: 4,
	},
}

var parseBoardingPassTestCases = []struct {
	input    string
	expected int
}{
	{
		input:    "BFFFBBFRRR",
		expected: 567,
	},
	{
		input:    "FFFBBBFRRR",
		expected: 119,
	},
	{
		input:    "BBFFBBFRLL",
		expected: 820,
	},
}
