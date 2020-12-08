package luggage

import "day-7/internal/luggage"

var parseRuleTestCases = []struct {
	input    string
	expected luggage.Rule
	success  bool
}{
	{
		input: "plaid lime bags contain 5 posh purple bags",
		expected: luggage.Rule{
			BagName: "plaid lime",
			ContainedBags: []*luggage.Bag{
				{
					Name: "posh purple",
					Qty:  5,
				},
			},
		},
		success: true,
	},
	{
		input: "dim violet bags contain no other bags",
		expected: luggage.Rule{
			BagName:       "dim violet",
			ContainedBags: nil,
		},
		success: true,
	},
	{
		input: "drab gold bags contain 2 posh aqua bags, 3 dark olive bags",
		expected: luggage.Rule{
			BagName: "drab gold",
			ContainedBags: []*luggage.Bag{
				{
					Name: "posh aqua",
					Qty:  2,
				},
				{
					Name: "dark olive",
					Qty:  3,
				},
			},
		},
		success: true,
	},
	{
		input: "dull beige bags contain 4 dim tan bags, 2 posh blue bags, 1 dull silver bag, 2 bright chartreuse bags",
		expected: luggage.Rule{
			BagName: "dull beige",
			ContainedBags: []*luggage.Bag{
				{
					Name: "dim tan",
					Qty:  4,
				},
				{
					Name: "posh blue",
					Qty:  2,
				},
				{
					Name: "dull silver",
					Qty:  1,
				},
				{
					Name: "bright chartreuse",
					Qty:  2,
				},
			},
		},
		success: true,
	},
}

var sumOuterBagsTestCases = []struct {
	input    []string
	expected int
	success  bool
}{
	{
		input: []string{
			"light red bags contain 1 bright white bag, 2 muted yellow bags",
			"dark orange bags contain 3 bright white bags, 4 muted yellow bags",
			"bright white bags contain 1 shiny gold bag",
			"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags",
			"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags",
			"dark olive bags contain 3 faded blue bags, 4 dotted black bags",
			"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags",
			"faded blue bags contain no other bags",
			"dotted black bags contain no other bags",
		},
		expected: 4,
		success:  true,
	},
}

var sumInnerBagsTestCases = []struct {
	input    []string
	expected int
	success  bool
}{
	{
		input: []string{
			"shiny gold bags contain 2 dark red bags",
			"dark red bags contain 2 dark orange bags",
			"dark orange bags contain 2 dark yellow bags",
			"dark yellow bags contain 2 dark green bags",
			"dark green bags contain 2 dark blue bags",
			"dark blue bags contain 2 dark violet bags",
			"dark violet bags contain no other bags",
		},
		expected: 126,
		success:  true,
	},
}
