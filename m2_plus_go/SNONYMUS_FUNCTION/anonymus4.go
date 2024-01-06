package main

import "fmt"

/*
Higher Order Functions in Golang
A Higher-Order function is a function that receives a function as an argument or returns the function as output.

Higher order functions are functions that operate on other functions, either by taking them as arguments or by returning them.
*/

//Passing Functions as Arguments to other Functions

func sum(x, y int) int {
	return x + y
}
func partialSum(x int) func(int) int {
	return func(y int) int {
		return sum(x, y)
	}
}
func main() {
	partial := partialSum(3)
	fmt.Println(partial(7))
}

/*
Output
10
*/
//In the program above, the partialSum function returns a sum function that takes two int arguments and returns a int argument.
