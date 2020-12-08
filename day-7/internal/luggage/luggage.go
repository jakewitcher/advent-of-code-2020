package luggage

import (
	"strconv"
	"strings"
)

type Bag struct {
	Name string
	Qty  int
}

type Rule struct {
	BagName       string
	ContainedBags []*Bag
}

type Rules map[string]*Rule

type Set map[string]interface{}

func SumOuterBags(input []string, target string) (int, error) {
	rules, err := ParseRules(input)
	if err != nil {
		return 0, err
	}

	outerBags := OuterBagsContainingTarget(target, rules)
	return len(outerBags), nil
}

func SumInnerBags(input []string, target string) (int, error) {
	rules, err := ParseRules(input)
	if err != nil {
		return 0, err
	}

	innerBags := QtyOfBagsContainedByTarget(target, rules)
	return innerBags, nil
}

func QtyOfBagsContainedByTarget(target string, rules Rules) int {
	var sum int

	targetBag := rules[target]
	for _, bag := range targetBag.ContainedBags {
		sum += bag.Qty + bag.Qty*QtyOfBagsContainedByTarget(bag.Name, rules)
	}

	return sum
}

func OuterBagsContainingTarget(target string, rules Rules) Set {
	outerBags := make(Set)
	seen := make(Set)

	outerBags[target] = struct{}{}
	var count int

	for count != len(outerBags) {
		count = len(outerBags)
		for bagName, _ := range outerBags {
			if _, ok := seen[bagName]; !ok {
				seen[bagName] = struct{}{}
				AddBagsContainingTarget(bagName, rules, outerBags)
			}
		}
	}

	delete(outerBags, target)
	return outerBags
}

func AddBagsContainingTarget(target string, rules Rules, bags Set) {
	for name, rule := range rules {
		for _, containedBag := range rule.ContainedBags {
			if containedBag.Name == target {
				bags[name] = struct{}{}
			}
		}
	}
}

func ParseRules(inputs []string) (Rules, error) {
	rules := make(map[string]*Rule)

	for _, input := range inputs {
		rule, err := ParseRule(input)
		if err != nil {
			return nil, err
		}

		rules[rule.BagName] = rule
	}

	return rules, nil
}

func ParseRule(input string) (*Rule, error) {
	words := strings.Split(input, " ")
	bagName := strings.Join(words[0:2], " ")

	containedBags, err := ParseContainedBags(strings.Join(words[4:], " "))
	if err != nil {
		return nil, err
	}

	rule := &Rule{
		BagName:       bagName,
		ContainedBags: containedBags,
	}

	return rule, nil
}

func ParseContainedBags(input string) ([]*Bag, error) {
	containedBagsDescriptions := strings.Split(input, ", ")

	if strings.HasPrefix(containedBagsDescriptions[0], "no") {
		return nil, nil
	}

	containedBags := make([]*Bag, len(containedBagsDescriptions))

	for i, bag := range containedBagsDescriptions {
		if strings.HasSuffix(bag, ",") {
			bag = strings.TrimSuffix(bag, ",")
		}

		words := strings.Split(bag, " ")
		name := strings.Join(words[1:3], " ")

		qty, err := strconv.Atoi(words[0])
		if err != nil {
			return nil, err
		}

		bag := &Bag{
			Name: name,
			Qty:  qty,
		}

		containedBags[i] = bag
	}
	return containedBags, nil
}
