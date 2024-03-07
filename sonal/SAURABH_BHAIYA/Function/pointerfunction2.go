package main

import "fmt"

func Add(a *int, b *int) {
	*a = 12
	*b = 13
	sum := *a + *b
	fmt.Println(sum)
	SUB(&a, &b)
}
func SUB(a **int, b **int) {
	**a = 15
	**b = 12
	Sub := **a - **b
	fmt.Println(Sub)
}
func main() {
	var a, b int
	Add(&a, &b)
}
