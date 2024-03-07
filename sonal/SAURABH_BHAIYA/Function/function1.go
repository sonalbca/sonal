package main

import "fmt"

//no return with no argument
//return with no argument
//no return with argument
//return with argument

func Addition(a *int, b *int) { //return with argumen
	*a = 5
	*b = 4
	fmt.Println("A :=", &a, "B :=", &b)

	ABC(&a, &b)
}
func ABC(A **int, B **int) {
	**A = 12
	**B = 15
}
func main() {
	var a, b int
	a = 6
	b = 7
	fmt.Println(&a, &b)
	fmt.Println(a, b)
	Addition(&a, &b)
	//fmt.Println(a, b)
	fmt.Println(&a, &b)
}
