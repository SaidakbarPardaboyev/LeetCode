func pacificAtlantic(heights [][]int) [][]int {
	rows := len(heights)
	cols := len(heights[0])

	res := [][]int{}
	visited := map[int]bool{}
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	memoizations := map[int][2]bool{}

	var dfs func(i, j int) [2]bool
	dfs = func(i, j int) [2]bool {
		checkerOceansFlow := [2]bool{}
		key := i*cols + j
		if _, ok := memoizations[key]; ok {
			checkerOceansFlow[0] = checkerOceansFlow[0] || memoizations[key][0]
			checkerOceansFlow[1] = checkerOceansFlow[1] || memoizations[key][1]
		}

		if i == 0 || j == 0 {
			checkerOceansFlow[0] = true
		}
		if i == rows-1 || j == cols-1 {
			checkerOceansFlow[1] = true
		}

		if checkerOceansFlow[0] && checkerOceansFlow[1] {
			return checkerOceansFlow
		}

		visited[key] = true
		for _, direction := range directions {
			r := direction[0] + i
			c := direction[1] + j
			if min(r, c) >= 0 && r < rows && c < cols &&
				heights[i][j] >= heights[r][c] && !visited[r*cols+c] {
				tem := dfs(r, c)
				checkerOceansFlow[0] = checkerOceansFlow[0] || tem[0]
				checkerOceansFlow[1] = checkerOceansFlow[1] || tem[1]
			}
		}
		delete(visited, key)

		memoizations[key] = checkerOceansFlow
		return checkerOceansFlow
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			tem := dfs(i, j)
			if tem[0] && tem[1] {
				res = append(res, []int{i, j})
			}
		}
	}

	return res
}