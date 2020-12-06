package form

func EvaluateTotalYesResponsesByGroup(group []string) int {
	yesResponses := make(map[rune]interface{})

	for _, memberAnswers := range group {
		for _, answer := range memberAnswers {
			yesResponses[answer] = struct{}{}
		}
	}

	return len(yesResponses)
}

func EvaluateConsensusYesResponsesByGroup(group []string) int {
	yesResponses := make(map[rune]int)
	members := len(group)

	for _, memberAnswers := range group {
		for _, answer := range memberAnswers {
			yesResponses[answer] += 1
		}
	}

	var sum int
	for _, responses := range yesResponses {
		if responses == members {
			sum++
		}
	}

	return sum
}
