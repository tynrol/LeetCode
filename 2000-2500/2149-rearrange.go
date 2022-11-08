package _000_2500

func rearrangeArray(nums []int) []int {
	res := make([]int, len(nums))
	pos, neg := 0, 1
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			res[pos] = nums[i]
			pos += 2
		} else {
			res[neg] = nums[i]
			neg += 2
		}
	}
	return res
}
