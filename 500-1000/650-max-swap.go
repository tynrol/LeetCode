package _00_1000

import "sort"

func maximumSwap(num int) int {
	list := []int{}
	start := num

	for num != 0 {
		list = append(list, num%10)
		num /= 10
	}

	list2 := make([]int, len(list))
	copy(list2, list)
	sort.Ints(list2)

	i := len(list) - 1
	for i > 0 {
		if list[i] == list2[i] {
			i--
		} else {
			for j := 0; j < i; j++ {
				if list[j] == list2[i] {
					list[i], list[j] = list[j], list[i]
					return out(list)
				}
			}
		}
	}
	return start
}

func out(list []int) int {
	res := 0
	for i := len(list) - 1; i >= 0; i-- {
		res = res*10 + list[i]
	}
	return res
}
