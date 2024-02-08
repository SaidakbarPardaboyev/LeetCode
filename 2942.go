package main

import (
	"fmt"
	"strings"
)

func main() {
	arr := []string{"leet", "code"}
	fmt.Println(findWordsContaining(arr, 'e'))
}

func findWordsContaining(words []string, x byte) []int {
	var res []int
	for i := 0; i < len(words); i++ {
		if strings.Contains(words[i], string(x)) {
			res = append(res, i)
		}
	}
	return res
}
