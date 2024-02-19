package main

import "fmt"

func TestFun(arr *[5]int) {
	fmt.Println("Array elements: ")
	for i := 0; i <= 4; i++ {
		fmt.Printf("%d ", (*arr)[i])
	}
}

func main() {
	arr := [...]int{1, 2, 3, 4, 5}
	TestFun(&arr)
}
