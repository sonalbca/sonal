package main

import "fmt"

func main() {
	SumOfArrayItems()
}

func SumOfArrayItems() {
	var size int
	sum := 0
	fmt.Println("enter the size of array")
	fmt.Scanln(&size)
	arr := make([]int, size)
	fmt.Println("enter the element of array")
	for i := 0; i < size; i++ {
		fmt.Scanln(&arr[i])
		sum = sum + arr[i]
	}
	fmt.Println("total sum of array is ", sum)
}
