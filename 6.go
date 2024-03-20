package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	ls := make([]string, numRows)
	lamp := true
	index := 0
	for i := 0; i < len(s); i++ {
		ls[index] += string(s[i])
		if lamp {
			index++
		} else {
			index--
		}
		if index == 0 {
			lamp = true
		} else if index == numRows-1 {
			lamp = false
		}
	}
	return strings.Join(ls, "")
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}
