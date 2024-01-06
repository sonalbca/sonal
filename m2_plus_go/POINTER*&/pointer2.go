package main

import "fmt"

func main() {
	var num int = 10
	var ptr *int
	var pptr **int

	ptr = &num
	pptr = &ptr

	fmt.Printf("Num: %d\n", num)
	fmt.Printf("*Ptr: %d\n", *ptr)
	fmt.Printf("**PPtr: %d\n", **pptr)

	**pptr = 20

	fmt.Printf("Num: %d\n", num)
	fmt.Printf("*Ptr: %d\n", *ptr)
	fmt.Printf("**PPtr: %d\n", **pptr)
}
