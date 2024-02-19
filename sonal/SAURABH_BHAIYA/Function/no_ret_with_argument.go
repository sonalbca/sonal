package main

import "fmt"

// function with Argument and no return
func sum(a int, b int) {
	sum := a + b
	fmt.Println(sum)
}
func main() {
	sum(12, 15)
}
