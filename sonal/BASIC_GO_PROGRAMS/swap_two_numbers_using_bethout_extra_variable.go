package main

import "fmt"

func SWap() {
	var a, b int
	a = 5
	b = 7
	fmt.Print(a, b)
	fmt.Println()
	a = a ^ b
	b = a ^ b
	fmt.Print(a, b)
}
func SHOP() {
	var item int
	fmt.Println("enter your item")
	fmt.Scanln(&item)
	var detol, park_avenue, lifeboy int
	detol = 45
	park_avenue = 67
	lifeboy = 35
	if item == detol {
		fmt.Println("yes in stock detol is present")
	}
	if item == park_avenue {
		fmt.Println("yes park is also present in stock")
	}
	if item == lifeboy {
		fmt.Println("life boy sabun to bahut hi koi leta hi nahi hai,lena hai")
	} else {
		fmt.Println("sorry invalid choice")
	}
}
func main() {
	SWap()
	//SHOP()
}
