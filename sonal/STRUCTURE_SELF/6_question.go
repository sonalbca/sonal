package main

import "fmt"

type Shape1 struct {
	area      float32
	perimeter float32
}

func (T Shape1) TRiangle() {
	var base, height, length, breath float32
	base = 2.1
	height = 3.2
	length = 4.2
	breath = 3.3
	T.area = (base * height) / 2
	fmt.Println("area of triangle is :=", T.area)
	fmt.Println()
	T.perimeter = 2 * (length + breath)
	fmt.Println("area of perimeter is :=", T.perimeter)
}
func main() {
	T := Shape1{}
	T.TRiangle()
}
