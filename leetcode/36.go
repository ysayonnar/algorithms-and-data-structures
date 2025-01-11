package leetcode

func is3x3Valid(matrix [3][3]byte) bool {
	set := make(map[int]bool)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			if matrix[i][j] != '.' {
				if _, ok := set[int(matrix[i][j])]; ok {
					return false
				} else {
					set[int(matrix[i][j])] = true
				}
			}
		}
	}

	return true
}

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		set := make(map[int]bool, len(board))
		for j := 0; j < len(board); j++ {
			if board[i][j] != '.' {
				if _, ok := set[int(board[i][j])]; ok {
					return false
				} else {
					set[int(board[i][j])] = true
				}
			}
		}
	}

	for j := 0; j < len(board); j++ {
		set := make(map[int]bool, len(board))
		for i := 0; i < len(board); i++ {
			if board[i][j] != '.' {
				if _, ok := set[int(board[i][j])]; ok {
					return false
				} else {
					set[int(board[i][j])] = true
				}
			}
		}
	}

	for row := 0; row < 9; row += 3 {
		for col := 0; col < 9; col += 3 {
			var square [3][3]byte

			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					square[i][j] = board[row+i][col+j]
				}
			}
			if !is3x3Valid(square) {
				return false
			}
		}
	}

	return true
}
