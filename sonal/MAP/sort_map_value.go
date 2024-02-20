package main

import (
	"fmt"
	"sort"
)

func main() {
	var SortInteger = map[string]int{"mango": 60, "guava": 20, "banana": 12, "plum": 30}

	values := make([]int, 0, len(SortInteger))

	for _, v := range SortInteger {
		values = append(values, v)
	}
	sort.Ints(values)

	for _, v := range values {
		fmt.Println(v)
	}
}
