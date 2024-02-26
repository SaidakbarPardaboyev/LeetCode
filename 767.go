package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reorganizeString("aaaaabbcdefgh"))
	fmt.Println(reorganizeString("abbabbaaab"))
	fmt.Println(reorganizeString("gpneqthatplqrofqgwwfmhzxjddhyupnluzkkysofgqawjyrwhfgdpkhiqgkpupgdeonipvptkfqluytogoljiaexrnxckeofqojltdjuujcnjdjohqbrzzzznymyrbbcjjmacdqyhpwtcmmlpjbqictcvjgswqyqcjcribfmyajsodsqicwallszoqkxjsoskxxstdeavavnqnrjelsxxlermaxmlgqaaeuvneovumneazaegtlztlxhihpqbajjwjujyorhldxxbdocklrklgvnoubegjrfrscigsemporrjkiyncugkksedfpuiqzbmwdaagqlxivxawccavcrtelscbewrqaxvhknxpyzdzjuhvoizxkcxuxllbkyyygtqdngpffvdvtivnbnlsurzroxyxcevsojbhjhujqxenhlvlgzcsibcxwomfpyevumljanfpjpyhsqxxnaewknpnuhpeffdvtyjqvvyzjeoctivqwann"))
}

func reorganizeString(s string) string {
	dic := MakeDic(s)
	res := MakeAString(dic)
	return res
}

func MakeAString(dic map[string]int) string {
	ls := SortMap(dic)
	var res []string
	for i := 0; i < len(ls[0][1]); i++ {
		res = append(res, string(ls[0][0]))
	}
	j := 1
	for i := 1; i < len(ls); i++ {
		for k := 0; k < len(ls[i][1]); k++ {
			if j <= len(res) {
				res = append(res[:j], append([]string{string(ls[i][0])}, res[j:]...)...)
			} else {
				j = 1
				res = append(res[:j], append([]string{string(ls[i][0])}, res[j:]...)...)
			}
			j += 2
		}
	}

	for i := 0; i < len(res)-1; i++ {
		if res[i] == res[i+1] {
			return ""
		}
	}
	return strings.Join(res, "")
}

func SortMap(dic map[string]int) [][]string {
	var ls [][]string
	for key, val := range dic {
		ls = append(ls, []string{key, strings.Repeat(key, val)})
	}

	for i := 0; i < len(ls); i++ {
		ind := i
		for j := i + 1; j < len(ls); j++ {
			if len(ls[ind][1]) < len(ls[j][1]) {
				ind = j
			}
		}
		ls[ind], ls[i] = ls[i], ls[ind]
	}
	return ls
}

func MakeDic(s string) map[string]int {
	dic := map[string]int{}
	for i := range s {
		if _, ok := dic[string(s[i])]; ok {
			dic[string(s[i])] += 1
		} else {
			dic[string(s[i])] = 1
		}
	}
	return dic
}
