package main

import "fmt"

func PrintNumber(num int) {
	if num <= 100 {
		fmt.Print(num, "\t\n")
		PrintNumber(num + 1)
	}
}

func main() {
	num := 1
	PrintNumber(num)
}
