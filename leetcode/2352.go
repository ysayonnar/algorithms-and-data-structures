package leetcode

func isArrayEqual(arr1, arr2 []int) bool {
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func EqualPairs(grid [][]int) int {
	cols := make(map[int][][]int)
	counter := 0

	for i := 0; i < len(grid); i++ {
		col := make([]int, len(grid))

		for j := 0; j < len(grid); j++ {
			col[j] = grid[j][i]
		}

		cols[col[0]] = append(cols[col[0]], col)
	}

	for i := 0; i < len(grid); i++ {
		if col, ok := cols[grid[i][0]]; ok {
			for _, item := range col {
				if isArrayEqual(grid[i], item) {
					counter++
				}
			}
		}
	}

	return counter
}
