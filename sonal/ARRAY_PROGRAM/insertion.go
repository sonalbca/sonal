package main

import "fmt"

func main() {

	// initializing array
	array := make([]int, 0, 8)
	array = append(array, 11, 20, 13, 44, 56, 96, 54, 97)
	fmt.Println("The given array is:", array)

	// getting the index
	var index = 4
	array[index] = 65
	fmt.Println()
	fmt.Println("The new array obtained after changing the element at", index, "index is:", array)
}
