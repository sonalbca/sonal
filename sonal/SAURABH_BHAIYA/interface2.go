package main

import "fmt"

type Employee interface {
	PrintName(name string)
	PrintSalary(basic int, tax int) int
}

func (e emp1) PrintName(name string) {
	fmt.Println("employee id", e)
	fmt.Println("employee name", name)
}
func (e emp1) PrintSalary(basic int, tax int) int {
	salary := (basic * tax) / 100
	return salary
}

type emp1 int

func main() {
	var e Employee
	e = emp1(1)
	e.PrintName("sonal")
	fmt.Println("employee salary", e.PrintSalary(2500, 15))

}
