package game

var testCases = []struct {
	input []int
	target int
	expected int
}{
	{
		input: []int{0,3,6},
		target: 2020,
		expected: 436,
	},
	{
		input: []int{1,3,2},
		target: 2020,
		expected: 1,
	},
	{
		input: []int{2,1,3},
		target: 2020,
		expected: 10,
	},
	{
		input: []int{1,2,3},
		target: 2020,
		expected: 27,
	},
	{
		input: []int{2,3,1},
		target: 2020,
		expected: 78,
	},
	{
		input: []int{3,2,1},
		target: 2020,
		expected: 438,
	},
	{
		input: []int{3,1,2},
		target: 2020,
		expected: 1836,
	},
	{
		input: []int{0,3,6},
		target: 30_000_000,
		expected: 175594,
	},
	{
		input: []int{1,3,2},
		target: 30_000_000,
		expected: 2578,
	},
	{
		input: []int{2,1,3},
		target: 30_000_000,
		expected: 3544142,
	},
	{
		input: []int{1,2,3},
		target: 30_000_000,
		expected: 261214,
	},
	{
		input: []int{2,3,1},
		target: 30_000_000,
		expected: 6895259,
	},
	{
		input: []int{3,2,1},
		target: 30_000_000,
		expected: 18,
	},
	{
		input: []int{3,1,2},
		target: 30_000_000,
		expected: 362,
	},
}
