package xmas

import (
	"day-9/internal/intslice"
	"errors"
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

	for i, j := 0, 2; j < len(sequence); i, j = i+1, i+3 {
		var sum int
		for sum < n || j > len(sequence) {
			rng := sequence[i:j]
			sum = intslice.Sum(rng)

			if sum == n {
				min, max := intslice.Min(rng), intslice.Max(rng)
				return max + min, nil
			}

			j++
		}
	}

	return 0, errors.New("no value found meeting condition")
}
