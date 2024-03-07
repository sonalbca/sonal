package main

import "fmt"

func printarr(array []int, size int) {
	for i := 0; i < size; i++ {
		fmt.Print(array[i], " ")
	}
	fmt.Println()
}
func lish(array []int, size int, toFind int) int {
	for i := 0; i < size; i++ {
		if array[i] == toFind {
			return i
		}
	}
	return -1
}
func main() {

	array := []int{10, 5, 3, 7, 6, 12}
	var toSearch int
	toSearch = 3
	printarr(array, 6)
	index := lish(array, 6, toSearch)
	if index == -1 {
		fmt.Println(toSearch, "is not present in the array.")
	} else {
		fmt.Println(toSearch, "is present at index ", index, "  in the array.")
	}
}
