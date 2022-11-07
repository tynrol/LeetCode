package main

import "strconv"

func isPalindrome(x int) bool {
	r := []rune(strconv.Itoa(x))
	i, j := 0, len(r)-1
	for i < j {
		if r[i] != r[j] {
			return false
		}
		i++
		j--
	}
	return true
}
