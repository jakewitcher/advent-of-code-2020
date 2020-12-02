package report

var twoEntriesTestCases = []struct {
	expenses []int
	target   int
	expected int
	success  bool
}{
	{
		expenses: []int{3, 4},
		target:   7,
		expected: 12,
		success:  true,
	},
	{
		expenses: []int{3, 4, 5, 6},
		target:   10,
		expected: 24,
		success:  true,
	},
	{
		expenses: []int{3, 4, 5, 6},
		target:   20,
		expected: 0,
		success:  false,
	},
	{
		expenses: []int{1721, 979, 366, 299, 675, 1456},
		target:   2020,
		expected: 514579,
		success:  true,
	},
}

var threeEntriesTestCases = []struct {
	expenses []int
	target   int
	expected int
	success  bool
}{
	{
		expenses: []int{3, 4, 5},
		target:   12,
		expected: 60,
		success:  true,
	},
	{
		expenses: []int{3, 4, 5, 6, 7, 8},
		target:   20,
		expected: 280,
		success:  true,
	},
	{
		expenses: []int{3, 4, 5, 6, 7, 8},
		target:   25,
		expected: 0,
		success:  false,
	},
	{
		expenses: []int{1721, 979, 366, 299, 675, 1456},
		target:   2020,
		expected: 241861950,
		success:  true,
	},
}