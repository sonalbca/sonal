package main

import "fmt"

// Go Program to Count Digits in a Number

func CountDigit(num int) int {
	count := 0
	for num != 0 {
		num = num / 10
		count++

	}
	fmt.Println(count)
	return 0
}
func main() {
	var num int
	fmt.Println("enter your given value")
	fmt.Scanln(&num)
	CountDigit(num)
}
