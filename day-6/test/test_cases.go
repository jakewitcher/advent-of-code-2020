package form

var testCases = []struct {
	input     []string
	total     int
	consensus int
}{
	{
		input: []string{
			"abc",
		},
		total:     3,
		consensus: 3,
	},
	{
		input: []string{
			"a",
			"b",
			"c",
		},
		total:     3,
		consensus: 0,
	},
	{
		input: []string{
			"ab",
			"ac",
		},
		total:     3,
		consensus: 1,
	},
	{
		input: []string{
			"a",
			"a",
			"a",
			"a",
		},
		total:     1,
		consensus: 1,
	},
	{
		input: []string{
			"b",
		},
		total:     1,
		consensus: 1,
	},
}
