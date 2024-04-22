func countSubstrings(s string, c byte) int64 {
	numberOfSubarray := 0

	lenOfLestSubarray := 0
	for _, letter := range s {
		if letter == rune(c) {
			lenOfLestSubarray++
			numberOfSubarray += lenOfLestSubarray
		}
	}
	return int64(numberOfSubarray)
}