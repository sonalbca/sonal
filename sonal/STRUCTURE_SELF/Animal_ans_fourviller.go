package main

import "fmt"

type Animal struct {
	name        string
	pet_animal  string
	wild_animal string
	eat         string
}
type Fourviller struct {
	NameOfCar string
	ComName   string
	ride      string
	tyre      int
	sits      int
}

func (a Animal) animal() {
	fmt.Println("Animal name is :=", a.name)
	fmt.Println("this animal is pet_animal :=", a.pet_animal)
	fmt.Println("this animal is wild_animal :=", a.wild_animal)
	fmt.Println("this animal eats :=", a.eat)
}
func (b Fourviller) fourviller() {
	fmt.Println("Name of car is :=", b.NameOfCar)
	fmt.Println("Name of company is :=", b.ComName)
	fmt.Println("the vehicle ride from :=", b.ride)
	fmt.Println("the tyre include is :=", b.tyre)
	fmt.Println("maximum sits available is :=", b.sits)

}
func main() {
	A := Animal{}
	A.name = "cow"
	A.pet_animal = "yes"
	A.wild_animal = "no"
	A.eat = "grass"
	A.animal()
	fmt.Println()

	B := Animal{}
	B.name = "lion"
	B.pet_animal = "no"
	B.wild_animal = "yes"
	B.eat = "meat"
	B.animal()
	fmt.Println()

	C := Animal{}
	C.name = "dog"
	C.pet_animal = "yes"
	C.wild_animal = "no"
	C.eat = "meat"
	C.animal()
	fmt.Println()

	D := Animal{}
	D.name = "tiger"
	D.pet_animal = "no"
	D.wild_animal = "yes"
	D.eat = "meat"
	D.animal()
	fmt.Println()

	a := Fourviller{}
	a.NameOfCar = "scorpio"
	a.ComName = "mahindra"
	a.ride = "diesel"
	a.tyre = 4
	a.sits = 7
	a.fourviller()
	fmt.Println()

	b := Fourviller{}
	b.NameOfCar = "Hexa"
	b.ComName = "tata"
	b.ride = "diesel"
	b.tyre = 4
	b.sits = 7
	b.fourviller()
	fmt.Println()

	c := Fourviller{}
	c.NameOfCar = "verna"
	c.ComName = "Hundai"
	c.ride = "diesel"
	c.tyre = 4
	c.sits = 5
	c.fourviller()
	fmt.Println()

	d := Fourviller{}
	d.NameOfCar = "nexon"
	d.ComName = "tata"
	d.ride = "diesel"
	d.tyre = 4
	d.sits = 5
	d.fourviller()

}
