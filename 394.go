package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func decodeString(s string) string {
	stack := []rune{}

	for i := 0; i < len(s); i++ {
		if s[i] != ']' {
			stack = append(stack, rune(s[i]))
		} else {
			str := ""
			for unicode.IsLetter(stack[len(stack)-1]) {
				str = string(stack[len(stack)-1]) + str
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]

			num := ""
			for len(stack) > 0 && unicode.IsDigit(stack[len(stack)-1]) {
				num = string(stack[len(stack)-1]) + num
				stack = stack[:len(stack)-1]
			}
			res, _ := strconv.Atoi(num)
			tem := strings.Repeat(str, res)

			for _, val := range tem {
				stack = append(stack, val)
			}
		}
	}
	return string(stack)
}

func main() {
	fmt.Println(decodeString("3[a2[c]]"))
}

// func decodeString(s string) string {
// 	stack := []interface{}{}
// 	res := ""
// 	num := 0
// 	var letter string
// 	i := 0
// 	for i < len(s) {
// 		if unicode.IsDigit(rune(s[i])) {
// 			num = int(s[i]) - '0'
// 			var j int
// 			for j = i + 1; j < len(s); j++ {
// 				if unicode.IsDigit(rune(s[j])) {
// 					num = num*10 + int(s[j]) - '0'
// 				} else {
// 					break
// 				}
// 			}
// 			stack = append(stack, num)
// 			i = j
// 			num = 0
// 		} else if unicode.IsLetter(rune(s[i])) {
// 			letter = string(s[i])
// 			var j int
// 			for j = i + 1; j < len(s); j++ {
// 				if unicode.IsLetter(rune(s[j])) {
// 					letter += string(s[j])
// 				} else {
// 					break
// 				}
// 			}
// 			stack = append(stack, letter)
// 			i = j
// 			letter = ""
// 		} else if s[i] == '[' {
// 			stack = append(stack, '[')
// 			i++
// 		} else if s[i] == ']' {
// 			st := ""
// 			for len(stack) > 0 {
// 				if str, ok := stack[len(stack)-1].(string); ok {
// 					s := str
// 					st = s + st
// 					stack = stack[:len(stack)-1]
// 				} else if stack[len(stack)-1] == '[' {
// 					stack = stack[:len(stack)-1]
// 					break
// 				}
// 			}
// 			put := strings.Repeat(st, stack[len(stack)-1].(int))
// 			stack = stack[:len(stack)-1]
// 			stack = append(stack, put)
// 			i++
// 		} else {
// 			i++
// 		}
// 	}
// 	for _, v := range stack {
// 		res += v.(string)
// 	}

// 	return res
// }
