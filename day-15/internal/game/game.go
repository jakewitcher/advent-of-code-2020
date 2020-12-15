package game

func FindNthNumberSpoken(input []int, n int) int {
	count := 1
	seen := make(map[int]int)

	for _, n := range input {
		seen[n] = count
		count++
	}

	var current int

	for count < n {
		var next int
		if lastSeen, ok := seen[current]; ok {
			next = count - lastSeen
		}

		seen[current] = count
		current = next
		count++
	}

	return current
}
