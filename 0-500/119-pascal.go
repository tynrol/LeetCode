package __500

func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1
	for i := 1; i <= rowIndex; i++ {
		val := row[i-1]
		val = val * (rowIndex + 1 - i)
		val = val / i
		row[i] = val
	}
	return row
}
