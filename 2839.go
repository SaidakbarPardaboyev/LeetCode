package main

import "fmt"

func main() {
	fmt.Println(canBeEqual("abcd", "cdab"))
}

func canBeEqual(s1 string, s2 string) bool {
	k := 0
	for k < len(s1) {
		for i := 0; i < len(s1)/2; i++ {
			s1 = s1[:i] + string(s1[i+2]) + string(s1[i+1]) + string(s1[i]) + s1[i+3:]
			if s1 == s2 {
				return true
			}
		}
		k++
	}
	return false
}
