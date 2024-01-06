package main

import "fmt"

func main() {

	var val = 0
XYZ:

	fmt.Println("Hello World")
	val = val + 1
	if val < 5 {
		goto XYZ
	}
}
