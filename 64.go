import "math"

func minPathSum(grid [][]int) int {
	directions := [][]int{{0, -1}, {-1, 0}}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// minimum number between the up index and left index
			minimumTilCurrentIndex := math.MaxInt64
			// To find minimum number between the up index and left index
			for _, directions := range directions {
				x := i + directions[0]
				y := j + directions[1]
				if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[i]) {
					minimumTilCurrentIndex = min(minimumTilCurrentIndex, grid[x][y])
				}
			}

			if minimumTilCurrentIndex != math.MaxInt64 {
				grid[i][j] += minimumTilCurrentIndex
			}
		}
	}

	return grid[len(grid)-1][len(grid[0])-1]
}