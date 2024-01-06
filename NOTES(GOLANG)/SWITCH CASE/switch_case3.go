package main

import "fmt"

func main() {
	var choice int
	var num1, num2 int
	fmt.Println("######## MUNU #######")
	fmt.Println("1st for addition")
	fmt.Println("2nd for subtractiob")
	fmt.Println("3rd for multiplication")
	fmt.Println("4th for division")
	fmt.Println("5th for module")
	fmt.Println()
	fmt.Println("enter your first number")
	fmt.Scanln(&num1)
	fmt.Println("enter your second number")
	fmt.Scanln(&num2)
	fmt.Println("enter your choice")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		result := num1 + num2
		fmt.Println(result)
	case 2:
		result := num1 - num2
		fmt.Println(result)
	case 3:
		result := num1 * num2
		fmt.Println(result)
	case 4:
		result := num1 % num2
		fmt.Println(result)
	case 5:
		result := num1 / num2
		fmt.Println(result)
	case 6:
		fmt.Println("invalid choice")
	}
}
