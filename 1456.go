package main

import (
	"fmt"
	"math"
)

func maxVowels(s string, k int) int {
	var res float64
	var numVowels float64
	res = 0
	numVowels = 0
	start := 0

	for end := 0; end < len(s); end++ {
		if s[end] == 'a' || s[end] == 'e' || s[end] == 'i' || s[end] == 'o' || s[end] == 'u' {
			numVowels++
		}

		if end >= k-1 {
			res = math.Max(res, numVowels)
			if s[start] == 'a' || s[start] == 'e' || s[start] == 'i' || s[start] == 'o' || s[start] == 'u' {
				numVowels--
			}
			start++
		}
	}
	return int(res)
}

func main() {
	fmt.Println(maxVowels("abciiidef", 3))
	fmt.Println(maxVowels("leetcode", 3))
}
