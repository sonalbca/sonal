package main

import "fmt"

// check weather a number is positive negative or zero using switch case
func main() {
	var num, choice int
	fmt.Println("enter any number")
	fmt.Scanln(&num)
	fmt.Println("enter your choice")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		if num > 0 {
			fmt.Println("positive number", num)
		}
		fmt.Println("not a positive number")
	case 2:
		if num < 0 {
			fmt.Println("negative number", num)
		}
		fmt.Println("not a negative number")
	default:
		fmt.Println("invalid choice")
	}
}
