func largestAltitude(gain []int) int {
	res := 0
	summa := 0
	for i := range gain {
		summa += gain[i]
		res = max(res, summa)
	}
	return res
}