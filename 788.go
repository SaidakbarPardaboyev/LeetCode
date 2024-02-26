package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(rotatedDigits(20))
	fmt.Println(rotatedDigits(9999))
}

func rotatedDigits(n int) int {
	res := 0
	for i := 2; i <= n; i++ {
		if Check(i) {
			res += 1
		}
	}
	return res
}

func Check(i int) bool {
	num := strconv.Itoa(i)
	ch := 0
	s := 0
	for i := range num {
		tem := string(num[i])
		if tem == "1" || tem == "0" || tem == "8" {
			s++
		} else if tem == "2" || tem == "5" || tem == "6" || tem == "9" {
			ch++
		} else {
			return false
		}
	}
	return s < len(num)
}
