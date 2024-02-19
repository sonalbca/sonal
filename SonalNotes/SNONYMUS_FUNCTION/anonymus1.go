package main

import "fmt"

func main() {
	Anonymus1()
	Anonymus2()
}

//Passing arguments to anonymous functions.

func Anonymus1() {
	func(l int, b int) {
		fmt.Println(l * b)
	}(20, 30)
}

//Function defined to accept a parameter and return value.

func Anonymus2() {
	fmt.Printf(
		"100 (°F) = %.2f (°C)\n",
		func(f float64) float64 {
			return (f - 32.0) * (5.0 / 9.0)
		}(100),
	)
}

//Anonymous functions can be used for containing functionality that need not be named and possibly for short-term use.
