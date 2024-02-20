package main

import "fmt"

func LinearSearch(Arr []int, size int, tofind int) int {
	for i := 0; i < size; i++ {
		if Arr[i] == tofind {
			return i
		}
	}
	return -1
}
func main() {
	Arr := []int{2, 5, 3, 8, 9, 6}
	var search = 9
	index := LinearSearch(Arr, 6, search)
	fmt.Println()
	if index == -1 {
		fmt.Println(search, "is not found")
	} else {
		fmt.Println(search, "is found", index, "number index")
	}
}
