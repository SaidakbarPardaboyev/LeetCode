package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	if len(chars) < 2 {
		return len(chars)
	}

	res := ""
	count := 1
	tem := ""
	for i := 1; i < len(chars); i++ {
		if chars[i] == chars[i-1] {
			count++
			if i == len(chars)-1 {
				tem += string(chars[i])
				if count > 1 {
					tem += strconv.Itoa(count)
				}
				res += tem
			}
		} else {
			tem += string(chars[i-1])
			if count > 1 {
				tem += strconv.Itoa(count)
			}
			if i == len(chars)-1 {
				tem += string(chars[i])
			}
			res += tem
			count = 1
			tem = ""
		}
	}

	for i := 1; i < len(res); i++ {
		chars[i] = res[i]
	}

	fmt.Println(res)
	return len(res)
}

func main() {
	fmt.Println(compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}))
	fmt.Println(compress([]byte{'a', 'b', 'c'}))
}
