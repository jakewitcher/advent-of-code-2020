package intslice

import "math"

func Sum(nums []int) int {
	var sum int
	for _, n := range nums {
		sum += n
	}

	return sum
}

func Min(nums []int) int {
	min := math.MaxInt64
	for _, n := range nums {
		if n < min {
			min = n
		}
	}

	return min
}

func Max(nums []int) int {
	max := math.MinInt64
	for _, n := range nums {
		if n > max {
			max = n
		}
	}

	return max
}
