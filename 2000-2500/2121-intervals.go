package _000_2500

func getDistances(arr []int) []int64 {
	pos := make(map[int][]int)
	sum := make(map[int]int)

	res := make([]int64, len(arr))

	for idx, num := range arr {
		pos[num] = append(pos[num], idx)
		sum[num] += idx
	}
	for num, idxs := range pos {
		var l, r int = 0, 0
		for i, idx := range idxs {
			r = sum[num] - l - idx
			res[idx] = int64(r - l - (len(idxs)-i-1-i)*idx)
			l += idx
		}
	}
	return res
}
