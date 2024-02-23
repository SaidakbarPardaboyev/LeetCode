package main

import (
	"fmt"
	"strings"
)

type MapSum struct {
	data map[string]int
}

func Constructor() MapSum {
	return MapSum{data: make(map[string]int)}
}

func (this *MapSum) Insert(key string, val int) {
	this.data[key] = val
}

func (this *MapSum) Sum(prefix string) int {
	sum := 0
	for key, val := range this.data {
		if strings.HasPrefix(key, prefix) {
			sum += val
		}
	}
	return sum
}

func main() {
	MapSum := Constructor()
	MapSum.Insert("apple", 3)
	fmt.Println(MapSum.Sum("ap"))
	MapSum.Insert("app", 2)
	fmt.Println(MapSum.Sum("ap"))
}
