func solve(board [][]byte) {
	lengthOfRow := len(board[0])
	lengthOfCol := len(board)

	visited := map[int]bool{}
	res := true
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i == lengthOfCol-1 || j == lengthOfRow-1 || i == 0 || j == 0 {
			res = false
			return
		}

		visited[i*lengthOfRow+j] = true
		for _, direction := range directions {
			r := direction[0] + i
			c := direction[1] + j
			if r >= 0 && r < lengthOfCol &&
				c >= 0 && c < lengthOfRow &&
				board[r][c] == 'O' && !visited[r*lengthOfRow+c] {
				dfs(r, c)
			}
		}
	}

	for i := 0; i < lengthOfCol; i++ {
		for j := 0; j < lengthOfRow; j++ {
			if board[i][j] == 'O' {
				dfs(i, j)
				if res {
					for val, _ := range visited {
						board[val/lengthOfRow][val%lengthOfRow] = 'X'
					}
				}
				visited = map[int]bool{}
				res = true
			}
		}
	}
}