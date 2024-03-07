package main

import "fmt"

type GetArea struct {
	area float64
}

func (G GetArea) setDim(width float64, breath float64) {
	G.area = width * breath
	fmt.Print("area of rectangle is :=", G.area)
}

func main() {
	G := GetArea{}
	G.setDim(12, 2.2)

}
