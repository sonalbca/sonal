package main

import "fmt"

// for Trianle
type triangle struct {
	height float64
	base   float64
}

// for recatangel
type Rectangle struct {
	width, height float64
}
type Circle struct {
	radius float64
}
type Shape struct {
	area      float64
	peremeter float64
}

func main() {
	t := triangle{height: 160.2, base: 160.1}
	r := Rectangle{width: 160.2, height: 160.2}
	c := Circle{radius: 5}

	// triagle area find
	s := Shape{}
	s.area = (t.height * t.base) / 2
	s.peremeter = 2 * t.height * t.base
	fmt.Println("Triangle: Area : ", s.area, " Peremetry : ", s.peremeter)

	// For Rectangle
	sR := Shape{}
	sR.area = r.height * r.width
	sR.peremeter = 2 * (r.height + r.width)
	fmt.Println("Rectangle Area : ", sR.area, " Peremetry : ", sR.peremeter)

	// For Circle
	Sc := Shape{}
	Sc.area = 3.14 * c.radius * c.radius
	Sc.peremeter = 2 * 3.14 * c.radius
	fmt.Println("Circle Area : ", Sc.area, " Circumference : ", Sc.peremeter)
}
