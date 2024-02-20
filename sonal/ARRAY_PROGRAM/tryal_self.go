package main

import (
	"fmt"
)

func ArithmeticOperation() {
	var a, b int
	a = 15
	b = 6
	fmt.Println("enter first number")
	fmt.Scanln(&a)
	fmt.Println("enter second number")
	fmt.Scanln(&b)
	a += b
	fmt.Println("sum of two numbers is ", a)
	a -= b
	fmt.Println("subtraction of two numbers is ", a)
	a *= b
	fmt.Println("multiplication of two number is ", a)
	a /= b
	fmt.Println("division of two number is ", a)
	a %= b
	fmt.Println("modulus of two number is ", a)
}
func INcrement() {
	var a = 5
	a++
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	fmt.Println(a)
}
func Decrement() {
	var a = 5
	a--
	for i := 10; i >= 1; i-- {
		fmt.Println(i)
	}
	fmt.Println(a)
}
func sssss() {
	var a, b int
	a = 3
	b = 4
	a += b
	b = +a
	fmt.Println(a)
	fmt.Println(b)
}
func ap() {
	var a, b int
	a = 15
	b = 201
	fmt.Println(a | b)
}
func Swap(a, b int) {
	fmt.Print("Before swapping,numbers are %d and %d\n", a, b)
	b = a + b
	a = b - a
	b = b - a
	fmt.Print("after swaping,numbers are %d and %d\n", a, b)
}
func main() {
	//ArithmeticOperation()
	//INcrement()
	//Decrement()
	//sssss()
	//ap()
	Swap(20, 10)
	Swap(3, 12)

	for row := 0; row < 7; row++ {
		if row == 0 || row == 5 {
			for col := 0; col < 7; col++ {
				fmt.Print("4")
			}
		} else {
			fmt.Print("4")
		}
		fmt.Println()
		for sp := 0; sp < 4; sp++ {
			fmt.Print("3")
		}
		fmt.Print("3")

	}
	fmt.Println()
}
