func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	} else if len(matrix[0]) == 1 && len(matrix) == 1 {
		return 1
	}

	memoization := map[int]int{}

	rows := len(matrix)
	cols := len(matrix[0])
	visited := map[int]bool{}
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	res := 0
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if _, ok := memoization[i*cols+j]; ok {
			return memoization[i*cols+j]
		}

		nestedMax := 1
		visited[i*cols+j] = true
		for _, direction := range directions {
			r := direction[0] + i
			c := direction[1] + j
			if r >= 0 && r < rows && c >= 0 && c < cols &&
				matrix[i][j] < matrix[r][c] && !visited[r*cols+c] {
				nestedMax = max(nestedMax, dfs(r, c)+1)
			}
		}
		delete(visited, i*cols+j)
		memoization[i*cols+j] = nestedMax
		res = max(res, nestedMax)
		return nestedMax
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			dfs(i, j)
			visited = map[int]bool{}
		}
	}

	return res
}