func lengthOfLongestSubstring(s string) int {
	res := 0
	cur := ""
	start := 0
	for end := 0; end < len(s); end++ {
		cur += string(s[end])
		if Check(cur) {
			res = max(res, end-start+1)
		} else {
			cur = cur[1:]
			start++
		}
	}
	return res
}

func Check(s string) bool {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}
	return true
}