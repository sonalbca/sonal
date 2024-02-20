package main

import "fmt"

func main() {
	var i, size int
	fmt.Println("enter the size of array")
	fmt.Scanln(&size)
	arr := make([]int, size)
	fmt.Println("enter the first array items")
	for i = 0; i < size; i++ {
		fmt.Scanln(&arr[i])
	}
	for i := 0; i < size; i++ {
		fmt.Print(arr[i])
	}
}
