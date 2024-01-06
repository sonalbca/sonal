package main

import "fmt"

func main() {
	var size, i int
	fmt.Println("enter the array size")
	fmt.Scanln(&size)
	arr := make([]int, size)
	arr1 := make([]int, size)
	arr2 := make([]int, size)
	fmt.Println("enter the first array items")
	for i = 0; i < size; i++ {
		fmt.Scanln(&arr[i])
	}
	fmt.Println("enter the second array items")
	for i = 0; i < size; i++ {
		fmt.Scanln(&arr1[i])
	}
	fmt.Println("the sum of addition of two array")
	for i = 0; i < size; i++ {
		arr2[i] = arr[i] + arr1[i]
		fmt.Println("", arr2[i])
	}

}
