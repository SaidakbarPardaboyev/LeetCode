func mergeAlternately(word1 string, word2 string) string {
	var res string
	var SmallLength int

	if len(word1) < len(word2) {
		SmallLength = len(word1)
	} else {
		SmallLength = len(word2)
	}

	i := 0
	for i < SmallLength {
		res += string(word1[i]) + string(word2[i])
		i++
	}

	if i < len(word1) {
		for i < len(word1) {
			res += string(word1[i])
			i++
		}
	}
	if i < len(word2) {
		for i < len(word2) {
			res += string(word2[i])
			i++
		}
	}
	return res
}
