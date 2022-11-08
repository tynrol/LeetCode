package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	len12 := len1 + len2
	last, curr := 0, 0
	first, second := 0, 0

	for first < len1 || second < len2 {
		last = curr
		if first < len1 && second < len2 {
			if nums1[first] < nums2[second] {
				curr = nums1[first]
				first++
			} else {
				curr = nums2[second]
				second++
			}
		} else if first < len1 {
			curr = nums1[first]
			first++
		} else if second < len2 {
			curr = nums2[second]
			second++
		}
		if float64(first+second) > float64(len12)/2 {
			break
		}
	}

	if len12%2 != 0 {
		return float64(curr)
	} else {
		return float64(last+curr) / 2
	}
}
