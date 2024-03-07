package main

import "fmt"

type Student struct {
	name string
	roll int
}

func main() {
	S := Student{}
	S.name = "john"
	S.roll = 2
	fmt.Println("Name is :=", S.name)
	fmt.Println("Roll no is :=", S.roll)
}
