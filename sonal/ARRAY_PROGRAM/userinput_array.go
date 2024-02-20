package main

import "fmt"

//Print Array Items

func main() {
	UserInput()
	//PrintArray()
}
func UserInput() {
	fmt.Println("enter size of your array")
	var size int
	fmt.Scanln(&size)
	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Printf("enter %dth element: ", i)
		fmt.Scanf("%d ", &arr[i])
	}
	fmt.Println("your array is :", arr)
}
func PrintArray() {
	var arraysize int
	fmt.Println("enter the size of array")
	fmt.Scanln(&arraysize)

	arr := make([]int, arraysize)
	for i, _ := range arr {
		fmt.Println("enter the array items", i)
		fmt.Scanln(&arr[i])
	}
	fmt.Println(arr)
}
