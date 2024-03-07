package main

import "fmt"

type Salary struct {
	Basic, HRA, TA float64
}
type employee struct {
	firstName, Lastname, Email string
	MonthelySalary             []Salary
}

func main() {
	fmt.Println("nested structure")
	e := employee{}
	e.firstName = "John"
	e.Lastname = "Doe"
	e.Email = "johndoe@gmail.com"
	e.MonthelySalary = []Salary{
		{Basic: 15000.0, HRA: 5000.0, TA: 2000.0},
		{Basic: 20001.0, HRA: 3000.0, TA: 1000.0},
	}

	fmt.Println(e)
}
