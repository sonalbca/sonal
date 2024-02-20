package main

import "fmt"

var number int

/*
var n, input int

	func main() {
		fmt.Println("enter any number")
		fmt.Scanln(&input)
		n = 1
		for n*2 <= input {
			n = n * 2
		}
		for n >= 1 {
			fmt.Print(input / n)
			input = input % n
			n = n / 2
		}
	}
*/
/*
func main() {
	fmt.Println("enter any number")
	fmt.Scanln(&number)
	var arr = [10]int{}
	i := 0
	for number > 0 {
		arr[i] = number % 2
		i++
		number = number / 2
	}
	for j := i - 1; j >= 0; j-- {
		fmt.Print (arr[j])
	}
}
*/
func main() { //binary to decimal
	var rem, num int
	fmt.Println("enter your binary number")
	fmt.Scanln(&num)
	for num != 0 {
		rem = num % 10
		num = num / 10
	}
}
