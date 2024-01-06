package main

import "fmt"

func main() {
	var num = 10
	var ptr *int

	ptr = &num
	fmt.Printf("Num: %d\n", num)
	fmt.Printf("*Ptr: %d\n", *ptr)

	*ptr = 20
	fmt.Printf("Num: %d\n", num)
	fmt.Printf("*Ptr: %d", *ptr)
}
