package main

import "fmt"

/*
      The return values of a function can be named in Golang
      Golang allows you to name the return values of a function.
	  We can also name the return value by defining variables,
      here a variable total of integer type is defined in the function declaration
      for the value that the function returns.
*/

func rectangle(l int, b int) (area int) {
	var parameter int
	parameter = 2 * (l + b)
	fmt.Println("Parameter: ", parameter)

	area = l * b
	return // Return statement without specify variable name
}

func main() {
	fmt.Println("Area: ", rectangle(20, 30))
}

/*
Output
Parameter:  100
Area:  600
*/
//Since the function is declared to return a value of type int,
//the last logical statement in the execution flow must be a return statement that returns a value of the declared type.
