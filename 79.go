func exist(board [][]byte, word string) bool {
	visited := map[int]bool{}
	length := len(board[0])
	res := false
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var backtracking func(i, j, ind int)

	backtracking = func(i, j, ind int) {
		if !(i >= 0 && i < len(board) &&
			j >= 0 && j < length && !visited[i*length+j]) {
			return
		}
		if word[ind] != board[i][j] {
			return
		}
		if ind == len(word)-1 {
			res = true
			return
		}

		visited[i*length+j] = true
		for _, dir := range directions {
			r := i + dir[0]
			c := j + dir[1]
			backtracking(r, c, ind+1)
		}
		delete(visited, i*length+j)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] != word[0] {
				continue
			}
			backtracking(i, j, 0)
		}
	}

	return res
}