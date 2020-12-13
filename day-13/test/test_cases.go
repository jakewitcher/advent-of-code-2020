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

var findMultipleOfXTestCases = []struct {
	x, y, z, offset int
	expected        int
}{
	//{
	//	x:        2,
	//	y:        7,
	//	z:        2,
	//	offset:   1,
	//	expected: 6,
	//},
	//{
	//	x:        4,
	//	y:        7,
	//	z:        4,
	//	offset:   1,
	//	expected: 20,
	//},
	//{
	//	x:        7,
	//	y:        13,
	//	z:        7,
	//	offset:   1,
	//	expected: 77,
	//},
	//{
	//	x:        67,
	//	y:        7,
	//	z:        402,
	//	offset:   1,
	//	expected: 804,
	//},
	//{
	//	x:        17,
	//	y:        13,
	//	z:        17,
	//	offset:   2,
	//	expected: 102,
	//},
	//{
	//	x:        17,
	//	y:        19,
	//	z:        17,
	//	offset:   3,
	//	expected: 187,
	//},
	//{
	//	x:        1789,
	//	y:        37,
	//	z:        1202161486,
	//	offset:   1,
	//	expected: 1202161486,
	//},
	//{
	//	x:        1789,
	//	y:        47,
	//	z:        1202161486,
	//	offset:   2,
	//	expected: 1202161486,
	//},
	//{
	//	x:        1789,
	//	y:        1889,
	//	z:        1202161486,
	//	offset:   3,
	//	expected: 1202161486,
	//},
	{
		x:        29,
		y:        37,
		z:        87870,
		offset:   101,
		expected: 88218,
	},
	{
		x:        29,
		y:        433,
		z:        87870,
		offset:   101,
		expected: 97324,
	},
	{
		x:        29,
		y:        13,
		z:        87870,
		offset:   101,
		expected: 87870,
	},
	{
		x:        29,
		y:        17,
		z:        87870,
		offset:   101,
		expected: 88044,
	},
	{
		x:        29,
		y:        19,
		z:        87870,
		offset:   101,
		expected: 88363,
	},
	{
		x:        29,
		y:        23,
		z:        87870,
		offset:   101,
		expected: 88334,
	},
	{
		x:        29,
		y:        977,
		z:        87870,
		offset:   101,
		expected: 105415,
	},
	{
		x:        29,
		y:        41,
		z:        87870,
		offset:   101,
		expected: 88131,
	},
}
