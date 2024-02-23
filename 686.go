package main

import "fmt"

func main() {
	fmt.Println(repeatedStringMatch("aaaaaaaaaaaaaaaaaaaaaab", "ba"))
	fmt.Println(repeatedStringMatch("a", "aa"))
	fmt.Println(repeatedStringMatch("aa", "a"))
}

func repeatedStringMatch(a string, b string) int {
	repeating := len(b) / len(a) * 3
	st := a
	if repeating == 0 {
		repeating = 2
	}
	for i := 1; i < repeating; i++ {
		st += string(a)
	}
	for i := 0; i < len(st); i++ {
		if st[i] == b[0] {
			if CheckSubset(st[i:], b) {
				return repeating - Cut(st[i+len(b):], a)
			}
		}
	}
	return -1
}

func Cut(st string, a string) int {
	res := 0
	for i := 0; i < len(st); {
		if st[i] == a[0] {
			if CheckSubset(st[i:], a) {
				res += 1
				i += len(a)
			} else {
				i++
			}
		} else {
			i++
		}
	}
	return res
}

func CheckSubset(st string, b string) bool {
	for i := 0; i < len(b); i++ {
		if !(i < len(st)) {
			return false
		}
		if st[i] != b[i] {
			return false
		}
	}
	return true
}
