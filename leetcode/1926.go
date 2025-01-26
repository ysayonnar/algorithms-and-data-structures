package leetcode

//я в рот ебал все что связано с графами
func nearestExit(maze [][]byte, entrance []int) int {
	queue := [][2]int{[2]int(entrance)}
	length := 0
	maze[entrance[0]][entrance[1]] = '+'

	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			n, m := queue[0][0], queue[0][1]
			queue = queue[1:]
			maze[n][m] = '+'
			if (n == 0 || n == len(maze)-1 || m == 0 || m == len(maze[0])-1) && (n != entrance[0] || m != entrance[1]) {
				return length
			}

			directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
			for _, dir := range directions {
				n1, m1 := n+dir[0], m+dir[1]
				if n1 >= 0 && m1 >= 0 && n1 < len(maze) && m1 < len(maze[0]) && maze[n1][m1] == '.' {
					queue = append(queue, [2]int{n1, m1})
					maze[n1][m1] = '+'
				}
			}
		}
		length++
	}

	return -1
}
