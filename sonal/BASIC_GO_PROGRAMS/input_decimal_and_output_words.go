package main

import "fmt"

func main() {
	var num int
	fmt.Println("enter the number")
	fmt.Scanln(&num)
	for num > 0 {
		rem := num % 10
		switch rem {
		case 1:
			fmt.Println("one")
			break
		case 2:
			fmt.Println("two")
			break
		case 3:
			fmt.Println("three")
			break
		case 4:
			fmt.Println("four")
			break
		case 5:
			fmt.Println("five")
			break
		case 6:
			fmt.Println("six")
			break
		case 7:
			fmt.Println("seven")
			break
		case 8:
			fmt.Println("eight")
			break
		case 9:
			fmt.Println("nine")
			break
		case 0:
			fmt.Println("zero")
			break
		default:
			fmt.Println("wrong")
			break
		}
		num = num / 10
	}
}
