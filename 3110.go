func scoreOfString(s string) int {
	sum := 0
	for i := 0; i < len(s)-1; i++ {
		sum += ABS(s[i], s[i+1])
	}
	return sum
}

func ABS(ch1, ch2 byte) int {
	res := int(ch1) - int(ch2)
	if res >= 0 {
		return res
	}
	return -res
}