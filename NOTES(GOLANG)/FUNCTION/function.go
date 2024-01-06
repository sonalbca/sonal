package main

import "fmt"

/*
What is Function in Golang
A function is a group of statements that exist within a program for the purpose of performing a specific task.
At a high level, a function takes an input and returns an output.

Function allows you to extract commonly used block of code into a single component.

The single most popular Go function is main(), which is used in every independent Go program.


Creating a Function in Golang
A declaration begins with the func keyword, followed by the name you want the function to have,
a pair of parentheses (), and then a block containing the function's code.

The following example has a function with the name SimpleFunction. It takes no parameter and returns no values.
*/

// SimpleFunction prints a message
func SimpleFunction() {
	fmt.Println("Hello World")
}

func main() {
	SimpleFunction()
}
