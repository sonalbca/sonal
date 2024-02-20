package main

import "fmt"

func main() {
	IntegerPalindrome()
	//StringPalindrome()
}

func IntegerPalindrome() {
	var num, rem, temp int
	sum := 0
	fmt.Println("Enter the number")
	fmt.Scanln(&num)
	temp = num
	for num > 0 {
		rem = num % 10
		sum = (sum * 10) + rem
		num = num / 10
	}
	if temp == sum {
		fmt.Println(temp, "Number is palindrome")
	} else {
		fmt.Println(temp, "Number is not palindrome")
	}
}
func StringPalindrome() {
	var name string
	fmt.Println("enter your name")
	fmt.Scanln(&name)
	rev := ""
	for _, i := range name {
		rev = string(i) + rev
	}
	fmt.Println(rev)
	if rev == name {
		fmt.Println(rev, "is palindrome")
	} else {
		fmt.Println(rev, "is not palindrome")
	}
}
