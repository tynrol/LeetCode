package Leetcode

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for idx, num := range nums {
		if req, is := numsMap[target-num]; is {
			return []int{req, idx}
		}
		numsMap[num] = idx
	}
	return []int{}
}
