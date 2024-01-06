package main

import "fmt"

func AddTwoInteger(num1 int, num2 int) {
	add := num1 + num2
	fmt.Println("sum of two numbers is ", add)
}
func main() {
	var num1, num2 int
	fmt.Println("enter first number")
	fmt.Scanln(&num1)
	fmt.Println("enter second number")
	fmt.Scanln(&num2)
	AddTwoInteger(num1, num2)
}
