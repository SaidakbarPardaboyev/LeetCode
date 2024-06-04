func longestPalindrome(s string) int {
	lettersWithCount := map[rune]int{}

	for _, chr := range s {
		lettersWithCount[chr]++
	}

	res := 0
	isThereOddNumber := false

	for _, count := range lettersWithCount {
		res += (count - count%2)
		if count%2 != 0 {
			if !isThereOddNumber {
				isThereOddNumber = true
				res++
			}
		}
	}

	return res
}