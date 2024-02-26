package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
}

func numJewelsInStones(jewels string, stones string) int {
	jewels = SortString(jewels)
	// fmt.Println(jewels)
	res := 0
	for i := range stones {
		if binarySearch(string(stones[i]), jewels) {
			res += 1
		}
	}
	return res
}

func binarySearch(el string, ls string) bool {
	start := 0
	finish := len(ls) - 1

	for start <= finish {
		mid := (start + finish) / 2

		tem := strings.Compare(string(ls[mid]), el)
		if tem == 0 {
			return true
		} else if tem == -1 {
			start = mid + 1
		} else {
			finish = mid - 1
		}
	}
	return false
}

func SortString(st string) string {
	var ls []string
	for i := range st {
		ls = append(ls, string(st[i]))
	}
	sort.Strings(ls)
	return strings.Join(ls, "")
}
