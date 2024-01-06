package main

import "fmt"

func main() {
	var choice int
	fmt.Println("########## MENU #########")
	fmt.Println("INDIA")
	fmt.Println("USA")
	fmt.Println("UK")
	fmt.Println("POK")
	fmt.Println()
	fmt.Println("enter your choice")
	fmt.Scanln(&choice)
	switch choice {
	case 5 - 1:
		fmt.Println("INDIA")
	case 5 - 2:
		fmt.Println("USA")
	case 5 - 3:
		fmt.Println("UK")
	case 5 - 4:
		fmt.Println("POK")
	case 5 - 5:
		fmt.Println("invalid choice")
	}
}
