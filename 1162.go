func maxDistance(grid [][]int) int {
    rows := len(grid)
    cols := len(grid[0])
    visited := map[int]bool{}
    directions := [][]int{{0,-1},{0,1},{1,0},{-1,0}}
    queue := []int{}

    var bfs func() int
    bfs = func() int {
        distance := -1
        for len(queue) > 0 {
            length := len(queue)
            for k := 0; k < length; k++ {
                i, j := queue[0]/cols, queue[0]%cols
                queue = queue[1:]
                for _, dir := range directions {
                    r, c := dir[0]+i, dir[1]+j
                    if min(r, c) >= 0 && r < rows && c < cols &&
                    grid[r][c] == 0 && !visited[r*cols+c] {
                        queue = append(queue, r*cols+c)
                        visited[r*cols+c] = true
                    }
                }
            }
            distance++
        }
        return distance
    }

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if grid[i][j] == 1 {
                queue = append(queue, i*cols+j)
            }
        }
    }

    if len(queue) == 0 || len(queue) == cols*rows {
        return -1
    }

    return bfs()
}