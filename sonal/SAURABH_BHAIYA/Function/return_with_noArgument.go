package main

import "fmt"

/*
// funcction with return but no argument
func abc() string {
	return " India"
}
func main() {
	st := "i love my"
	fmt.Print(st)
	fmt.Println(abc())
}
*/

// return with argument
func Arithmetic(a int, b int) int {
	sum := a + b
	sub := a - b
	mul := a * b
	div := a / b
	mod := a % b
	fmt.Println("sum of two number is ", sum)
	fmt.Println("subtraction of two number is ", sub)
	fmt.Println("multiplication of two number is ", mul)
	fmt.Println("division of two number is ", div)
	return mod
}
func main() {
	var a, b int
	fmt.Println("Enter first number")
	fmt.Scanln(&a)
	fmt.Println("Enter second number")
	fmt.Scanln(&b)
	res := Arithmetic(a, b)
	fmt.Println("modulus of two number is ", res)
}
