func replaceWords(dictionary []string, sentence string) string {
	words := map[byte][]string{}
	sen := strings.Split(sentence, " ")
	for _, word := range dictionary {
		words[word[0]] = append(words[word[0]], word)
	}

	res := ""
	fmt.Println(words)

	for _, word := range sen {
		curWord := word
		for _, dic := range words[word[0]] {
			l := len(dic)
			sgm := l
			if l <= len(word) && dic == word[:sgm] {
				if len(curWord) > sgm {
					curWord = word[:sgm]
				}
			}
		}
		res += curWord + " "
	}

	if len(res) == 0 {
		return res
	}

	return res[:len(res)-1]
}