package main

import (
	"strconv"
)

func digitSum(s string, k int) string {
	if len(s) <= k {
		return s
	}

	str := ""
	for i := 0; i < len(s); i += k {
		group, n := s[i:min(i+k, len(s))], 0

		for _, digit := range group {
			num, _ := strconv.Atoi(string(digit))
			n += num
		}
		str += strconv.Itoa(n)
	}
	s = str

	return digitSum(s, k)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
