func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	lenghtOfRow := len(obstacleGrid[0])
	lenghtOfCol := len(obstacleGrid)
	visited := map[int]int{}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i >= lenghtOfCol || j >= lenghtOfRow ||
			obstacleGrid[i][j] == 1 {
			return 0
		}

		if i == lenghtOfCol-1 && j == lenghtOfRow-1 {
			return 1
		}

		if _, ok := visited[i*lenghtOfRow+j]; ok {
			return visited[i*lenghtOfRow+j]
		}

		numberOfPath := dfs(i, j+1)
		numberOfPath += dfs(i+1, j)

		visited[i*lenghtOfRow+j] = numberOfPath

		return numberOfPath
	}

	return dfs(0, 0)
}