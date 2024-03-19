package main

import (
	"fmt"
	"sort"
)

func closeStrings(word1 string, word2 string) bool {
	dic1 := make(map[byte]int)
	dic2 := make(map[byte]int)

	for i := range word1 {
		dic1[word1[i]]++
	}
	for i := range word2 {
		dic2[word2[i]]++
	}

	if len(dic1) != len(dic2) {
		return false
	}

	var letters []byte
	var counts1 []int
	var counts2 []int

	for key, count := range dic1 {
		letters = append(letters, key)
		counts1 = append(counts1, count)
		counts2 = append(counts2, dic2[key])
	}

	if len(letters) != len(dic1) {
		return false
	}

	sort.Ints(counts1)
	sort.Ints(counts2)

	for i := 0; i < len(counts1); i++ {
		if counts1[i] != counts2[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(closeStrings("aabbcccddd", "abccccdddd"))
}
