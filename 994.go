func orangesRotting(grid [][]int) int {
    rows := len(grid)
    cols := len(grid[0])

    queue := [][]int{}
    freshs := 0
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if grid[i][j] == 2 {
                queue = append(queue, []int{i, j})
            }else if grid[i][j] == 1 {
                freshs++
            }
        }
    }

    if freshs == 0 {
        return 0
    }

    sikl := 0
    directions := [][]int{{0,1},{0,-1},{1,0},{-1,0}}
    for len(queue) > 0 {
        length := len(queue)
        for k := 0; k < length; k++ {
            i, j := queue[0][0], queue[0][1]
            queue = queue[1:]
            for _, dir := range directions {
                r := dir[0] + i
                c := dir[1] + j
                if min(r, c) >= 0 && r < rows && c < cols &&
                grid[r][c] == 1 {
                    grid[r][c] = 2
                    freshs--
                    queue = append(queue, []int{r, c})
                }
            }
        }
        sikl++
    }

    if freshs != 0{
        return -1
    }

    fmt.Println(grid)
    return sikl-1
}