package main

import "fmt"

func AddTwoNum(array []int, array1 []int, size int) []int {
	var i int
	arr := make([]int, size)

	for i = 0; i < len(array); i++ {
		fmt.Print(" ", array[i])
	}
	fmt.Println()
	for i = 0; i < len(array1); i++ {
		fmt.Print(" ", array1[i])
	}
	fmt.Println()
	for i = 0; i < len(arr); i++ {
		arr[i] = array[i] + array1[i]

	}
	return arr
}
func main() {
	array := []int{2, 4, 3}
	array2 := []int{5, 6, 4}
	arr := AddTwoNum(array, array2, 3)
	fmt.Print(" ", arr)

}
