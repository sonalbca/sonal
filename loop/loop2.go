package main

import "fmt"

func main() {
	var row int
	fmt.Println("enter row")
	fmt.Scanln(&row)
	for i := 1; i <= row; i++ {
		for sp := 1; sp < row-i; sp++ {
			fmt.Print(" ")
		}
		for col := i; col >= 1; col-- {
			fmt.Print(col)

		}
		for col := 2; col <= i; col++ {
			fmt.Print(col)
		}
		fmt.Println()
	}
}
