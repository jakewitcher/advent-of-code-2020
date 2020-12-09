package xmas

import (
	"errors"
	"math"
)

func FindFirstWeakness(sequence []int, preamble int) (int, error) {
Main:
	for i, j := 0, preamble; j < len(sequence); i, j = i+1, j+1 {
		rng := sequence[i:j]
		target := sequence[j]

		for n := 0; n < len(rng)-1; n++ {
			for m := n + 1; m < len(rng); m++ {
				if rng[n]+rng[m] == target {
					continue Main
				}
			}
		}

		return target, nil
	}

	return 0, errors.New("no value found meeting condition")
}

func FindSecondWeakness(sequence []int, preamble int) (int, error) {
	n, err := FindFirstWeakness(sequence, preamble)
	if err != nil {
		return 0, err
	}

	sum := sequence[0]

	i, j := 0, 1
	for j <= len(sequence) {
		if sum == n {
			rng := sequence[i:j]
			min, max := MinMax(rng)
			return min + max, nil
		}

		if sum > n {
			sum -= sequence[i]
			i++
		}

		if sum < n {
			sum += sequence[j]
			j++
		}
	}

	return 0, errors.New("no value found meeting condition")
}

func MinMax(nums []int) (min, max int) {
	min, max = math.MaxInt64, math.MinInt64
	for _, n := range nums {
		if n < min {
			min = n
		}

		if n > max {
			max = n
		}
	}

	return min, max
}
