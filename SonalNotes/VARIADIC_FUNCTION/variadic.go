package main

import "fmt"

/*
Variadic Functions
A variadic function is a function that accepts a variable number of arguments. In Golang, it is possible to pass a varying number of arguments of the same type as referenced in the function signature. To declare a variadic function, the type of the final parameter is preceded by an ellipsis, "...", which shows that the function may be called with any number of arguments of this type. This type of function is useful when you don't know the number of arguments you are passing to the function, the best example is built-in Println function of the fmt package which is a variadic function.

Select single argument from all arguments of variadic function
This program demonstrates the use of a variadic function in Golang. A variadic function is a function that takes a variable number of arguments of a specific type.

In this program, the function variadicExample takes a variadic parameter of type string, indicated by the ... before the type name. This means that the function can accept any number of string arguments.

In the main function, we call variadicExample with four string arguments: "red", "blue", "green", and "yellow". These arguments are passed to the s parameter of the variadicExample function as a slice of strings.
*/

func main() {
	VariadicExample("red", "blue", "green", "yellow")
}

func VariadicExample(s ...string) {
	fmt.Println(s[0])
	fmt.Println(s[3])
}

/*
Output
red
yellow
*/
//Because we are accessing the first and fourth elements of the s slice. Note that if we were to pass fewer than four arguments to variadicExample, the program would still run without error,
//but trying to access an index beyond the length of the slice would result in a runtime error.
