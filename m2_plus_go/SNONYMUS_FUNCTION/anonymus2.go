package main

import "fmt"

/*
Closures Functions in Golang
Closures are a special case of anonymous functions. Closures are anonymous functions which access the variables defined outside the body of the function.
*/

//Anonymous function accessing the variable defined outside body.

func main() {
	l := 20
	b := 30

	func() {
		var area int
		area = l * b
		fmt.Println(area)
	}()
}
