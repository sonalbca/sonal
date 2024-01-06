package main

import "fmt"

/*
Anonymous Functions in Golang
An anonymous function is a function that was declared without any named identifier to refer to it.
Anonymous functions can accept inputs and return outputs, just as standard functions do.

Assigning function to the variable.
*/

var (
	area = func(l int, b int) int {
		return l * b
	}
)

func main() {
	fmt.Println(area(20, 30))
}
