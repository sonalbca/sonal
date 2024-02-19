package main

import "fmt"

//User Defined Function Types in Golang
//Golang also support to define our own function types.

//The modified version of above program with function types as below:

type First func(int) int
type Second func(int) First

func SquareSum(x int) Second {
	return func(y int) First {
		return func(z int) int {
			return x*x + y*y + z*z
		}
	}
}

func main() {
	// 5*5 + 6*6 + 7*7
	fmt.Println(SquareSum(5)(6)(7))
}
