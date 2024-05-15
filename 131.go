func partition(s string) [][]string {
	res := [][]string{}
	var PalindromePartitioning func(ind int, curSub []string)

	PalindromePartitioning = func(ind int, curSub []string) {
		if ind == len(s) {
			res = append(res, append([]string{}, curSub...))
			return
		}

		for i := ind; i < len(s); i++ {
			if IsPalindrome(s[ind : i+1]) {
				curSub = append(curSub, s[ind:i+1])
				PalindromePartitioning(i+1, curSub)
				curSub = curSub[:len(curSub)-1]
			}
		}
	}

	PalindromePartitioning(0, []string{})

	return res
}

func IsPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i <= j; i++ {
		if s[i] != s[j] {
			return false
		}
		j--
	}
	return true
}