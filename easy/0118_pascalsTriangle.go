package easy

func generate(numRows int) [][]int {
	// initialize and make first row
	res := make([][]int, numRows)
	res[0] = []int{1}

	// quick return for small numRows
	if numRows == 1 {
		return res
	}

	res[1] = []int{1, 1}

	// fill out insides
	for i := 2; i < numRows; i++ {
		// make new row and fill out 1s at edges
		res[i] = make([]int, i+1)
		res[i][0], res[i][i] = 1, 1

		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}

	return res
}
