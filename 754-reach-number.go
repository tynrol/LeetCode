package Leetcode

import "math"

func reachNumber(target int) int {
	target = int(math.Abs(float64(target)))
	var sum, steps int = 0, 0

	for sum < target {
		steps++
		sum += steps
	}
	for (sum-target)%2 != 0 {
		steps++
		sum += steps
	}
	return steps
}
