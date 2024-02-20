package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(updateMatrix([][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}))
}

func updateMatrix(mat [][]int) [][]int {
	res := make([][]int, len(mat))
	for i := range res {
		res[i] = make([]int, len(mat[0]))
	}

	var ls [][]int
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 0 {
				ls = append(ls, []int{i, j})
			}
		}
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 0 {
				res[i][j] = 0
			} else {
				dist := math.MaxInt
				for k := range ls {
					temI := i - ls[k][0]
					temJ := j - ls[k][1]
					if temI < 0 {
						temI = -temI
					}
					if temJ < 0 {
						temJ = -temJ
					}
					if temI+temJ < dist {
						dist = temI + temJ
					}
				}
				res[i][j] = dist
			}
		}
	}

	return res
}
