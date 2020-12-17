package ticket

var parseFieldTestCaes = []struct {
	input string
	name  string
	set   []int
}{
	{
		input: "zone: 1-3 or 5-7",
		name:  "zone",
		set:   []int{1, 2, 3, 5, 6, 7},
	},
	{
		input: "departure time: 10-13 or 15-17",
		name:  "departure time",
		set:   []int{10, 11, 12, 13, 15, 16, 17},
	},
}
