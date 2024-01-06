package main

import "fmt"

/*
Golang Functions Returning Multiple Values
Functions in Golang can return multiple values, which is a helpful feature in many practical scenarios.

This example declares a function with two return values and calls it from a main function.
*/

func Rectangle(l int, b int) (area int, parameter int) {
	parameter = 2 * (l + b)
	area = l * b
	return // Return statement without specify variable name
}

func main() {
	var a, p int
	a, p = Rectangle(20, 30)
	fmt.Println("Area:", a)
	fmt.Println("Parameter:", p)
}
