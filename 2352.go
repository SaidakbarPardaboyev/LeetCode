func equalPairs(grid [][]int) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			lamp := true
			for k := 0; k < len(grid); k++ {
				if grid[i][k] != grid[k][j] {
					lamp = false
					break
				}
			}
			if lamp {
				res++
			}
		}
	}
	return res
}