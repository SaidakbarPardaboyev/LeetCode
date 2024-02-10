package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(lastVisitedIntegers([]string{"1", "2", "prev", "prev", "prev", "prev"}))
}

func lastVisitedIntegers(words []string) []int {
	var res []int
	ls := []int{}
	ind := 1
	for _, v := range words {
		tem, err := strconv.Atoi(v)
		if err != nil {
			if len(ls) < ind {
				res = append(res, -1)
			} else {
				res = append(res, ls[len(ls)-ind])
				ind += 1
			}
		} else {
			ls = append(ls, tem)
			ind = 1
		}
	}
	return res
}
