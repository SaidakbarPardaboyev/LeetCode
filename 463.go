func islandPerimeter(grid [][]int) int {
	perimetrOfIsland := 0

	sizeOfColumn := len(grid)
	for i := 0; i < sizeOfColumn; i++ {
		sizeOfRow := len(grid[i])
		for j := 0; j < sizeOfRow; j++ {
			if grid[i][j] == 1 {
				AddToPeremetrByCheckPositionOfCellInIsland(i, j, grid, &perimetrOfIsland)
			}
		}
	}
	return perimetrOfIsland
}

func AddToPeremetrByCheckPositionOfCellInIsland(x, y int, grid [][]int, perimetrOfIsland *int) {
	if x == 0 || grid[x-1][y] == 0 {
		*perimetrOfIsland++
	}
	if x == len(grid)-1 || grid[x+1][y] == 0 {
		*perimetrOfIsland++
	}
	if y == 0 || grid[x][y-1] == 0 {
		*perimetrOfIsland++
	}
	if y == len(grid[0])-1 || grid[x][y+1] == 0 {
		*perimetrOfIsland++
	}
}