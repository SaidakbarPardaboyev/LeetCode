package main

import (
	"fmt"
	"strings"
)

type stack []string

func (s *stack) Add(str byte) {
	*s = append(*s, string(str))
}

func (s *stack) Pop() {
	if len(*s) == 0 {
		return
	}
	*s = (*s)[:len(*s)-1]
}

func removeStars(s string) string {
	stack := stack{}
	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			stack.Pop()
		} else {
			stack.Add(s[i])
		}
	}
	return strings.Join(stack, "")
}

func main() {
	// ptr := stack{}
	// ptr.Add('s')
	// ptr.Add('a')
	// ptr.Add('l')
	// ptr.Add('o')
	// ptr.Add('m')
	// fmt.Println(ptr)

	fmt.Println(removeStars("leet**cod*e"))

}
