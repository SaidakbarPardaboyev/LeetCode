
func findMinimumOperations(s1 string, s2 string, s3 string) int {
	res := 0
	if s1 == s2 && s2 == s3 {
		return res
	}
	var lamp = false
	for len(s1) != 0 && len(s2) != 0 && len(s3) != 0 {
		if len(s1) >= len(s2) && len(s1) >= len(s3) {
			s1 = s1[:len(s1)-1]
		} else if len(s2) >= len(s1) && len(s2) >= len(s3) {
			s2 = s2[:len(s2)-1]
		} else if len(s3) >= len(s1) && len(s3) >= len(s2) {
			s3 = s3[:len(s3)-1]
		}
		res++
		if s1 == s2 && s2 == s3 {
			lamp = true
			break
		}
	}
	if lamp {
		return res
	}
	return -1
}