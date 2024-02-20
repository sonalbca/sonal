package main

import "fmt"

func AddEachDigitAndCount(num int) int {
	sum := 0
	count := 0
	for num != 0 {
		modulus := num % 10
		sum = sum + modulus
		count++
		num = num / 10
	}
	fmt.Println("sum of given value is ", sum)
	fmt.Println("count of each digit is ", count)
	return 0
}
func main() {
	var num int
	fmt.Println("enter your given value")
	fmt.Scanln(&num)
	AddEachDigitAndCount(num)
}
