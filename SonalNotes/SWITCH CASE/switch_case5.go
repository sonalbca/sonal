package main

//find maximum between two numbers using switch case
import "fmt"

func main() {
	var num1, num2, choice int
	fmt.Println("enter first number")
	fmt.Scanln(&num1)
	fmt.Println("enter second number")
	fmt.Scanln(&num2)
	fmt.Println("enter your choice")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		if num1 > num2 {
			fmt.Println("num1 is greater than num2", num1)
		} else if num2 > num1 {
			fmt.Println("num2 is greater than num1", num2)
		}
	default:
		fmt.Println("invalid choice")
	}
}
