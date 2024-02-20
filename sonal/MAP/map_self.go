package main

import (
	"fmt"
	"sort"
)

func main() {
	var unSorted = map[string]int{"Sonal": 83, "Harshit": 12, "Chintu": 89}
	keys := make([]string, 0, len(unSorted))
	for k, _ := range unSorted {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, v := range keys {
		fmt.Println(v, unSorted[v])
	}
}
