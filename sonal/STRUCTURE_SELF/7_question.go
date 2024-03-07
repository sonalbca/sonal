package main

import "fmt"

type shape struct {
	area      float32
	perimeter float32
}

func (s shape) Triangle1() {
	var base, height, length, width float32
	base = 3
	height = 4
	length = 3
	width = 4
	s.area = (base * height) / 2
	fmt.Println("area of triangle is ", s.area)
	s.perimeter = length * width
	fmt.Println("area of perimeter", s.perimeter)

}
func main() {
	s := shape{}
	s.Triangle1()

}
