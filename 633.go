func judgeSquareSum(c int) bool {
	start := 0
	end := int(math.Sqrt(float64(c)))

	for start <= end {
		if start*start+end*end == c {
			return true
		} else if start*start+end*end > c {
			end--
		} else {
			start++
		}
	}
	return false
}