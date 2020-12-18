package homework

var parseExpressionEqualPrecedenceTestCases = []struct {
	input    string
	expected int
}{
	{
		input:    "2 * 2 + 3",
		expected: 7,
	},
	{
		input:    "2 + 2 * 3",
		expected: 12,
	},
	{
		input:    "2 * 2 + 3 * 7",
		expected: 49,
	},
	{
		input:    "(1 + 2)",
		expected: 3,
	},
	{
		input:    "1 + (3 * 2)",
		expected: 7,
	},
	{
		input:    "(3 * (4 * 8) * 5 * 7 * 3) + 8",
		expected: 10088,
	},
	{
		input:    "4 * (6 * 4 * 6) + 8 + 4",
		expected: 588,
	},
}

var parseExpressionAdditionHigherPrecedenceTestCases = []struct {
	input    string
	expected int
}{
	{
		input:    "2 * 3 + 1",
		expected: 8,
	},
	{
		input:    "4 + 2 * 3 + 1",
		expected: 24,
	},
	{
		input:    "4 + (2 * 3) + 1",
		expected: 11,
	},
	{
		input:    "1 + (2 * 3) + (4 * (5 + 6))",
		expected: 51,
	},
	{
		input:    "2 * 3 + (4 * 5)",
		expected: 46,
	},
	{
		input:    "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
		expected: 669060,
	},
	{
		input:    "5 + (8 * 3 + 9 + 3 * 4 * 3)",
		expected: 1445,
	},
	{
		input:    "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
		expected: 23340,
	},
}
