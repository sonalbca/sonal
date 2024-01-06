package main

import "fmt"

/*
Passing multiple string arguments to a variadic function
This program demonstrates the use of variadic function in Go, which allows a function to accept a variable number of arguments of the same type.

In this example, the function variadicExample() is defined to accept a variadic parameter of type string. This means that it can accept any number of string arguments. The main() function calls variadicExample() multiple times with different numbers of string arguments.

The first call to variadicExample() is made without any arguments, which is allowed since the function is defined to accept zero or more string arguments.

The second, third, and fourth calls pass different numbers of string arguments to variadicExample(). In each case, the function prints the contents of the s parameter using fmt.Println().
*/

func main() {

	variadicExample()
	variadicExample("red", "blue")
	variadicExample("red", "blue", "green")
	variadicExample("red", "blue", "green", "yellow")
}

func variadicExample(s ...string) {
	fmt.Println(s)
}

/*
Output
[]
[red blue]
[red blue green]
[red blue green yellow]
*/

//The first call to variadicExample() prints an empty slice since no arguments were passed.
//The subsequent calls print the contents of the s parameter, which contains all the string arguments passed to the function.
