package main

import "fmt"

type Triangle struct {
	breath int
	height int
}
type Rictangle struct {
	length int
	width  int
}
type square struct {
	num int
}

func main() {
	a := Triangle{}
	a.breath = 12
	a.height = 24
	area := 1 / 2 * a.breath * a.height
	fmt.Println("area of triangle is ", area)
	R := Rictangle{}
	R.length = 12
	R.width = 24
	Ric := R.length * R.width
	fmt.Println(Ric)
	sq := square{}
	sq.num = 2
	squ := sq.num * sq.num
	fmt.Printf("square of %d is %d", sq.num, squ)
}
