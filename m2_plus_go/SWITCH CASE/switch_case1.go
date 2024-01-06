package main

import "fmt"

func main() {
	var choice int
	fmt.Println("########## MENU ###########")
	fmt.Println("Sunday")
	fmt.Println("monday")
	fmt.Println("tuesday")
	fmt.Println("wednesday")
	fmt.Println("thursday")
	fmt.Println("friday")
	fmt.Println("saturday")
	fmt.Println("enter your choice")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		fmt.Println("sunday")
	case 2:
		fmt.Println("monday")
	case 3:
		fmt.Println("tuesday")
	case 4:
		fmt.Println("wednesday")
	case 5:
		fmt.Println("thursday")
	case 6:
		fmt.Println("friday")
	case 7:
		fmt.Println("saturday")
	}

}
