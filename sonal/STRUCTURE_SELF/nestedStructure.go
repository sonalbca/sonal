package main

import "fmt"

// nested structure

type person struct {
	firstPerson string
	details     struct {
		age    int
		heigth float64
		dt     struct {
			gender string
		}
	}
}

func main() {
	// first instance
	p := person{}
	p.firstPerson = "James"
	p.details.heigth = 160.4
	p.details.age = 24
	p.details.dt.gender = "Male"
	fmt.Println(p)

	// Second Instance
	p2 := person{}
	p2.firstPerson = "Bond"
	p2.details.age = 26
	p2.details.heigth = 167.4
	p2.details.dt.gender = "Female"
	fmt.Println(p2)

}
