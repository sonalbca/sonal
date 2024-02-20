package main

import "fmt"

func Addition(a int, b int) int {
	total := a + b
	fmt.Println("sum of total number is :=", total)
	return 0
}
func main() {
	var a, b int
	fmt.Println("enter first number")
	fmt.Scanln(&a)
	fmt.Println("enter second number")
	fmt.Scanln(&b)
	Addition(a, b)
}
