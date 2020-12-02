package report

import "fmt"

func FindProductOfTwoEntriesWithSum(entries []int, target int) (int, error) {
	for i := 0; i < len(entries)-1; i++ {
		for j := i + 1; j < len(entries); j++ {
			if entries[i]+entries[j] == target {
				return entries[i] * entries[j], nil
			}
		}
	}

	return 0, fmt.Errorf("no pair of entries found equaling %v", target)
}

func FindProductOfThreeEntriesWithSum(entries []int, target int) (int, error) {
	for i := 0; i < len(entries)-2; i++ {
		for j := i + 1; j < len(entries)-1; j++ {
			for k := i + 2; k < len(entries); k++ {
				if entries[i]+entries[j]+entries[k] == target {
					return entries[i] * entries[j] * entries[k], nil
				}
			}
		}
	}

	return 0, fmt.Errorf("no pair of entries found equaling %v", target)
}