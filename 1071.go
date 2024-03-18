package main

import (
	"fmt"
	"strings"
)

func gcdOfStrings(str1 string, str2 string) string {
	var res string

	var small string
	var large string
	if len(str1) > len(str2) {
		small = str2
		large = str1
	} else {
		small = str1
		large = str2
	}
	var tem string
	for i := range small {
		tem += string(str1[i])
		if large == strings.Repeat(tem, len(large)/len(tem)) {
			if small == strings.Repeat(tem, len(small)/len(tem)) {
				res = tem
			}
		}
		if len(tem) >= len(small) {
			break
		}
	}

	return res
}

func main() {
	fmt.Println(gcdOfStrings("TAUXXTAUXXTAUXXTAUXXTAUXX", "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX"))
}
