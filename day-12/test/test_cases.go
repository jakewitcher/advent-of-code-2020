package navigation

var testCases = []struct {
	input            []string
	naiveDistance    int
	waypointDistance int
}{
	{
		input: []string{
			"F10",
			"N3",
			"F7",
			"R90",
			"F11",
		},
		naiveDistance:    25,
		waypointDistance: 286,
	},
}
