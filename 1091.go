func shortestPathBinaryMatrix(grid [][]int) int {
    rows := len(grid)
    cols := len(grid[0])
    directions := [][]int{{1,0},{-1,0},{0,1},{0,-1},{-1,-1},{-1,1},{1,-1},{1,1}}
    visited := map[int]bool{}
    var bfs func() int
    bfs = func() int {
        queue := []int{0}
        visited[0] = true
        count := 1
        for len(queue) > 0 {
            length := len(queue)
            for k := 0; k < length; k++ {
                i, j := queue[0]/cols, queue[0]%cols
                if i == rows-1 && j == cols-1 {
                    return count
                }
                queue = queue[1:]
                for _, dir := range directions {
                    r, c := dir[0]+i, dir[1]+j
                    if min(r,c) >= 0 && r < rows && c < cols &&
                    grid[r][c] == 0 && !visited[r*cols+c] {
                        queue = append(queue, r*cols+c)
                        visited[r*cols+c] = true
                    }
                }
            }
            count++
        }
        return -1
    }

    if grid[0][0] != 0 {
        return -1
    }

    return bfs()
}