func wordBreak(s string, wordDict []string) []string {
	res := []string{}

	words := map[string]bool{}
	for _, val := range wordDict {
		words[val] = true
	}
	var FindSubSet func(ind int, subSet string)
	FindSubSet = func(ind int, subSet string) {
		if ind == len(s) {
			res = append(res, subSet[:len(subSet)-1])
			return
		}

		for i := ind; i < len(s); i++ {
			if words[s[ind:i+1]] {
				FindSubSet(i+1, subSet+s[ind:i+1]+" ")
			}
		}
	}
	FindSubSet(0, "")
	return res
}