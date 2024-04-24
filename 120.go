func minimumTotal(triangle [][]int) int {
	// Update the numbers in the array with add smaller number
	// from the i and i + 1 index of the next row
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			if triangle[i+1][j] > triangle[i+1][j+1] {
				triangle[i][j] += triangle[i+1][j+1]
			} else {
				triangle[i][j] += triangle[i+1][j]
			}
		}
	}

	return triangle[0][0]
}