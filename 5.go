package main

import "fmt"

func longestPalindrome(s string) string {
	res := ""
	str := ""
	for start := 0; start < len(s); start++ {
		for end := start; end < len(s); end++ {
			str = s[start : end+1]
			if Check(str) {
				if len(str) > len(res) {
					res = str
				}
			}
		}
	}
	return res
}

func Check(str string) bool {
	for start, end := 0, len(str)-1; start < end; start++ {
		if str[start] != str[end] {
			return false
		}
		end--
	}
	return true
}

func main() {
	fmt.Println("res := ", longestPalindrome("cbbd"))
	fmt.Println("res := ", longestPalindrome("cabababa"))
	fmt.Println("res := ", longestPalindrome("babad"))
}
