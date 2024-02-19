package main

import "fmt"

// check weather a number is even or odd using swich case
func main() {
	var a, choice int
	fmt.Println("#########MENU#######")
	fmt.Println("1 for even number")
	fmt.Println("2 for odd number")

	fmt.Println("enter any number")
	fmt.Scanln(&a)
	fmt.Println("enter your choice")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		if a%2 == 0 {
			fmt.Println("even number", a)
		}
	case 2:
		if a%2 != 0 {
			fmt.Println("odd number", a)
		}
	default:
		fmt.Println("invalid choice")
	}
}
