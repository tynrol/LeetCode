package main

func isMonotonic(nums []int) bool {
	//I wanted here to sort a list and check if its different, but go isn't like pnums[i+1]thon
	if nums[0] < nums[len(nums)-1] {
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				return false
			}
		}
	} else {
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] < nums[i+1] {
				return false
			}
		}
	}
	return true
}

func isMonotonic2(nums []int) bool {
	isIncreasing := nums[0] > nums[len(nums)-1]
	for i := 0; i < len(nums)-1; i++ {
		if !(nums[i] == nums[i+1] || (nums[i] > nums[i+1] && isIncreasing) || (nums[i] < nums[i+1] && !isIncreasing)) {
			return false
		}
	}
	return true
}
