package main

import "fmt"

func main() {
	var num = 1.12
	var ptr *float64
	ptr = &num
	fmt.Println("Num is ", num)
	fmt.Println("*ptr is ", *ptr)

	*ptr = 12.32
	fmt.Println("Num is ", num)
	fmt.Println("*ptr is ", *ptr)
}
