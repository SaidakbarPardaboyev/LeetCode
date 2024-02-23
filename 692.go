package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2))
	fmt.Println(topKFrequent([]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4))

}

func topKFrequent(words []string, k int) []string {
	sorted_words := SortWords(words)
	var res []string
	for i := 0; i < k; i++ {
		res = append(res, sorted_words[i][0])
	}
	return res
}

func SortWords(words []string) [][]string {
	dic := map[string]int{}
	for i := range words {
		if _, ok := dic[words[i]]; ok {
			dic[words[i]] += 1
		} else {
			dic[words[i]] = 1
		}
	}
	var wordsLs [][]string
	for key, val := range dic {
		wordsLs = append(wordsLs, []string{key, strings.Repeat("a", val)})
	}
	Sort(wordsLs)
	return wordsLs
}

func Sort(words [][]string) {
	for i := 0; i < len(words); i++ {
		ind := i
		for j := i + 1; j < len(words); j++ {
			if words[ind][1] < words[j][1] {
				ind = j
			} else if words[ind][1] == words[j][1] {
				tem := strings.Compare(words[ind][0], words[j][0])
				if tem > 0 {
					ind = j
				}
			}
		}
		words[ind], words[i] = words[i], words[ind]
	}
}
