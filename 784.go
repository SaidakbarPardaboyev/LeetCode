func letterCasePermutation(s string) []string {

	var findLetterPermutation func(i int) []string
	findLetterPermutation = func(i int) []string {
		if i == len(s) {
			return []string{""}
		}

		resPer := []string{}
		curPer := findLetterPermutation(i + 1)
		for _, per := range curPer {
			resPer = append(resPer, string(s[i])+per)
			if s[i] >= 'A' && s[i] <= 'Z' {
				resPer = append(resPer, string(s[i]+32)+per)
			}
			if s[i] >= 'a' && s[i] <= 'z' {
				resPer = append(resPer, string(s[i]-32)+per)
			}
		}
		return resPer
	}
	return findLetterPermutation(0)
}