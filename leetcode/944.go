package leetcode

func countNormalOranges(grid [][]int) int {
	ctr := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				ctr++
			}
		}
	}
	return ctr
}

// ладно графы я дружу с вами
func orangesRotting(grid [][]int) int {
	queue := [][2]int{}
	m, n := len(grid), len(grid[0])

	normal := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			} else if grid[i][j] == 1 {
				normal++
			}
		}
	}
	if normal == 0 {
		return 0
	} else if len(queue) == 0 {
		return -1
	}

	minutes := 0
	for len(queue) > 0 {
		normalCtr := countNormalOranges(grid)
		if normalCtr == 0 {
			return minutes
		}

		l := len(queue)
		for i := 0; i < l; i++ {
			x, y := queue[0][0], queue[0][1]
			queue = queue[1:]
			grid[x][y] = 2

			directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
			for _, dir := range directions {
				x1, y1 := x+dir[0], y+dir[1]
				if x1 >= 0 && y1 >= 0 && x1 < m && y1 < n {
					if grid[x1][y1] == 1 {
						queue = append(queue, [2]int{x1, y1})
						grid[x1][y1] = 2
					}
				}
			}
		}
		minutes++
	}

	if countNormalOranges(grid) != 0 {
		return -1
	}

	return minutes
}
