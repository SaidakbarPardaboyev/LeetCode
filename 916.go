package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wordSubsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"e", "oo"}))
	fmt.Println(wordSubsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"l", "e"}))

}

func wordSubsets(words1 []string, words2 []string) []string {
	var res []string
	// Ikkinchi listdan unique belgilarni bitta string qilish
	word := MakeSett(words2)
	// Birinchi listdagi hammasini solishtirish
	for i := range words1 {
		// Solishtiruvchi funcsiyaga bervorish
		if CompareWords(words1[i], &word) {
			res = append(res, words1[i])
		}
	}
	return res
}

// Unique belgilarni va ularni bitta indexda eng ko'p qatnashgan sonini olish
func MakeSett(st []string) map[string]int {
	// Dictionary orqali unique belgilarni va max sonini olish
	dic := make(map[string]int)
	for _, str := range st {
		for _, letter := range str {
			if _, ok := dic[string(letter)]; !ok {
				dic[string(letter)] = strings.Count(str, string(letter))
			} else {
				tem := strings.Count(str, string(letter))
				if tem > dic[string(letter)] {
					dic[string(letter)] = tem
				}
			}
		}
	}
	return dic
}

// Solishtiruvchi funcsiyaga
func CompareWords(FirstWord string, word *map[string]int) bool {
	for letter, sch := range *word {
		if !(strings.Count(FirstWord, string(letter)) >= sch) {
			return false
		}
	}
	return true
}
