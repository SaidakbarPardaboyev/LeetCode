func numEnclaves(grid [][]int) int {
    rows := len(grid)
    cols := len(grid[0])

    directions := [][]int{{0,-1},{0,1},{1,0},{-1,0}}
    res := 0
    visited := map[int]bool{}

    var bfs func(i, j int) int
    bfs = func(i, j int) int {
        queue := []int{i*cols+j}
        isNonWalkableToBorder := false
        numberOfCell := 0

        for len(queue) > 0 {
            length := len(queue)
            for k := 0; k < length; k++ {
                i, j := queue[0]/cols, queue[0]%cols
                queue = queue[1:]

                if min(i, j) == 0 || i == rows-1 || j == cols-1 {
                    isNonWalkableToBorder = true
                } 

                visited[i*cols+j] = true
                for _, dir := range directions {
                    r := i + dir[0]
                    c := j + dir[1]
                    if min(r, c) >= 0 && r < rows && c < cols &&
                    grid[r][c] == 1 && !visited[r*cols+c] {
                        visited[r*cols+c] = true
                        queue = append(queue, r*cols+c)
                    }
                }
                numberOfCell++
            }
        }

        if !isNonWalkableToBorder {
            return numberOfCell
        }
        return 0
    }

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if grid[i][j] == 1 && !visited[i*cols+j]{
                numberOfCells := bfs(i, j)
                res += numberOfCells
            }
        }
    }

    fmt.Println(grid)

    return res
}