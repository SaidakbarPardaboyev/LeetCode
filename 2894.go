func differenceOfSums(n int, m int) int {
	sum1 := 0
	sum2 := 0
	for v := 1; v <= n; v++ {
		if v%m == 0 {
			sum2 += v
		} else {
			sum1 += v
		}
	}
	return sum1 - sum2
}
