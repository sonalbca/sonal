package main

import "fmt"

/*
Simple function with parameters in Golang
Information can be passed to functions through arguments. An argument is just like a variable.

Arguments are specified after the function name, inside the parentheses. You can add as many arguments as you want, just separate them with a comma.

The following example has a function with two arguments of int type. When the add() function is called, we pass two integer values (e.g. 20,30).
*/

// Function accepting arguments
func add(x int, y int) {
	total := 0
	total = x + y
	fmt.Println(total)
}

func main() {
	// Passing arguments
	add(20, 30)
}

/*
Output
50

If the functions with names that start with an uppercase letter will be exported to other packages.
If the function name starts with a lowercase letter, it won't be exported to other packages,
but you can call this function within the same package.
*/
