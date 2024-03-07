package main

import "fmt"

type Person struct {
	a int
	b float32
}

func main() {
	var c Person
	c.a = 5
	c.b = 3.2
	fmt.Println(c)
}
