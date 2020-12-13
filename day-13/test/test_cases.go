package shuttle

var nextAvailableBusTestCases = []struct {
	input     []string
	timestamp int
	expected  int
}{
	{
		input:     []string{"7", "13", "x", "x", "59", "x", "31", "19"},
		timestamp: 939,
		expected:  295,
	},
}

var earliestTimestampTestCases = []struct {
	input    []string
	expected int
}{
	{
		input:    []string{"7", "13", "x", "x", "59", "x", "31", "19"},
		expected: 1068781,
	},
	{
		input:    []string{"17", "x", "13", "19"},
		expected: 3417,
	},
	{
		input:    []string{"67", "7", "59", "61"},
		expected: 754018,
	},
	{
		input:    []string{"67", "x", "7", "59", "61"},
		expected: 779210,
	},
	{
		input:    []string{"67", "7", "x", "59", "61"},
		expected: 1261476,
	},
	{
		input:    []string{"1789", "37", "47", "1889"},
		expected: 1202161486,
	},
}

