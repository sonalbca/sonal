package main

import (
	"fmt"
	"math"
)

/*
(1) Golang Program to Print Hello World
(2) Go Program to add Two Numbers
(3) Go Program to Find the Compound Interest
(4) Go Program to Count Digits in a Number
(5) Go Program to Count Total Notes in an Amount
(6) Go Program to Find the Cube of a Number
(7) Go Program to Calculate the Employee Salary
//(8) Go Program to calculate Electricity Bill
(9) Go Program to check Even or Odd
(10) Go Program to print Even Numbers from 1 to N
(11) Go Program to find Factors of a Number
(12) Go Program to find Factorial of a Number
(13) Go Program to find First Digit of a Number
//(14) Go Program to find Generic Root of a Number
(15) Go Program to check Largest of Two Numbers
(16) Go Program to check Largest of Three Numbers
(17) Go Program to check Leap year
(18) Go Program to print Multiplication Table
(19) Go Program to print Natural Numbers from 1 to N
(20) Go Program to print Natural Numbers in Reverse Order
//(21) Go Program to find NCR Factorial of a Number
(22) Go Program to find Number divisible by 5 and 11
(23) Go Program to print Odd Numbers 1 to N
(24) Go Program to find Product of Digits in a Number
(25) Check Palindrome Number
(26) Check Perfect Number
(27) Check Prime Number
(28) Check Positive or Negative
(29) Calculate Power of a Number
(30) Calculate Profit or Loss
(31) Print 1 to 100
(32) Print 1 to 100 without using loop
(33) Reverse a Number
//(34) Roots of a Quadratic Equation
(35) Simple Interest
(36) Square of a Number
(37) Square root of a Number
(38) Sum and Average of Natural Numbers
(39) Sum of Digits in a Number//ye nahi bana hai
(40) Sum of Even Numbers
(41) Sum of Even and Odd Numbers
(42) Sum of Odd Numbers
(43) Swap Two Numbers
*/
func main() {
	//HelloWorld()
	//Addition()
	//Addition1()
	//CompoundInterest()//ye nahi bana
	CountDigit()
	//AddEachDigit() //add each digit and count
	//CountTotalNotes()
	//CubePro()
	//CalculateEmployeeSalary()
	//CheckEvenOdd()
	//EvenNumber1ToN()
	//Factorss()
	//Factorial1()
	//FirstDigit()//ye nahi bana
	//LargestTwoNum()
	//LargestThreeNum()
	//CheckLeapYear()
	//MultiplicationTable()
	//NaturalNumber1To10()
	//NaturalNumbersReverseOrder()
	//Divisibleby5and11()
	//PrintOddNumbers1toN()
	//ProductOfNumber()
	//PalindromeNumber()
	//PerfectNumber()
	//PrimeNumber()
	//PositiveorNegative()
	//PowerOfANumber()
	//ProfitOrLoss()
	//Print1To100()
	//ReverseNumber()
	//SimpleInterest()
	//SquarerootofaNumber()
	//SquareNum()
	//SumandAverageofNaturalNumbers()
	//SumofDigitsinaNumber()//ye nahi bana hai
	//SumOfEvenNumber()
	//SumOfEvenAndOddNumber()
	//SumOfOddNumber()
	//SwapTwoNumber()
	//SWAPTWONUMBERWITHOUTTEMPVARIABLE()
	//swapTwonumberusingbitwise()

}

func HelloWorld() {
	fmt.Println("hello world")

}
func Addition() {
	var num1, num2 int
	fmt.Println("enter first number")
	fmt.Scanf("%d", &num1)
	fmt.Println("enter second number")
	fmt.Scanf("%d", &num2)
	addition := num1 + num2
	fmt.Printf("addition of two numbers is %d\n", addition)

}
func Addition1() {
	num1 := 12
	num2 := 13
	addition := num1 + num2
	fmt.Printf("addition of two numbers is %d", addition)
}
func CompoundInterest() {
	var principal, rate, time int
	fmt.Println("enter principal")
	fmt.Scanln(&principal)
	fmt.Println("enter rate")
	fmt.Scanln(&rate)
	fmt.Println("enter time")
	fmt.Scanln(&time)
	principal = (principal * rate / 10) ^ time
	fmt.Println(principal)

}
func CountDigit() {
	count := 0
	var num int
	fmt.Println("enter any number")
	fmt.Scanf("%d", &num)
	for num != 0 {
		num = num / 10
		count++

	}
	fmt.Printf("%d", count)
}
func AddEachDigit() {
	count := 0
	sum := 0
	num := 0
	fmt.Println("enter any number")
	fmt.Scanln(&num)
	for num != 0 {
		remainder := num % 10
		sum = sum + remainder
		num = num / 10
		count++
	}
	fmt.Println(sum)
	fmt.Println(count)
}
func CountTotalNotes() {
	var amount int
	Notes := [8]int{2000, 500, 200, 50, 10, 5, 2, 1}
	fmt.Println("enter your amount")
	fmt.Scanln(&amount)
	temp := amount
	for i := 0; i < 7; i++ {
		fmt.Println(Notes[i], "Notes :=", temp/Notes[i])
		temp = temp % Notes[i]
	}
}
func CubePro() {
	var num int
	fmt.Println("enter any number")
	fmt.Scanln(&num)
	cube := num * num * num
	fmt.Println("Cube of num is :=", cube)

}
func CalculateEmployeeSalary() {

	var basicSal, hra, da, grossSal float64

	fmt.Print("Enter the Employee Basic Salary = ")
	fmt.Scanln(&basicSal)

	if basicSal <= 10000 {
		hra = (basicSal * 8) / 100
		da = (basicSal * 10) / 100
	} else if basicSal <= 20000 {
		hra = (basicSal * 16) / 100
		da = (basicSal * 20) / 100
	} else {
		hra = (basicSal * 24) / 100
		da = (basicSal * 30) / 100
	}

	grossSal = basicSal + hra + da
	fmt.Println("The Gross Salary of this Employee = ", grossSal)
}
func CheckEvenOdd() {
	var num int
	fmt.Println("enter any number")
	fmt.Scanln(&num)
	if num%2 == 0 {
		fmt.Printf("even number %d\n ", num)
	} else {
		fmt.Printf("odd number %d\n", num)
	}

}
func EvenNumber1ToN() {
	var num int
	fmt.Println("enter any number")
	fmt.Scanln(&num)
	for i := 1; i <= num; i++ {
		if i%2 == 0 {
			fmt.Println("even number is", i)
		}
	}
}
func Factorss() {
	var num, i int
	fmt.Println("enter any numbers")
	fmt.Scanln(&num)
	for i = 1; i <= num; i++ {
		if num%i == 0 {
			fmt.Println("factors of num is ", i)
		}
	}
}
func Factorial1() {
	var num int
	factorial := 1
	fmt.Println("enter any number")
	fmt.Scanln(&num)
	for i := 1; i <= num; i++ {
		factorial = factorial * i
	}
	fmt.Println(factorial)
}
func FirstDigit() {
	var num, rem int
	fmt.Println("enter first number")
	fmt.Scanf("%d", &num)
	for num >= 10 {
		rem = num / 10
	}
	fmt.Printf("%d", rem)
}
func LargestTwoNum() {
	var num1, num2 int
	fmt.Println("enter first number")
	fmt.Scanln(&num1)
	fmt.Println("enter second number")
	fmt.Scanln(&num2)
	if num1 > num2 {
		fmt.Println("greatest number is ", num1)
	} else {
		fmt.Println("greatest number is ", num2)
	}
}
func LargestThreeNum() {
	var num1, num2, num3 int
	fmt.Println("enter first number")
	fmt.Scanln(&num1)
	fmt.Println("enter second number")
	fmt.Scanln(&num2)
	fmt.Println("enter third number")
	fmt.Scanln(&num3)
	if num1 > num2 && num1 > num3 {
		fmt.Println("greatest number is ", num1)
	} else if num2 > num1 && num2 > num3 {
		fmt.Println("greatest number is ", num2)
	} else {
		fmt.Println("greatest number is ", num3)
	}
}
func CheckLeapYear() {
	var year int
	fmt.Println("enter your year")
	fmt.Scanln(&year)
	if year%4 == 0 {
		fmt.Println(year, "is leap year")
	} else {
		fmt.Println(year, "is not leap year")
	}
}
func MultiplicationTable() {
	var num int
	mul := 1
	fmt.Println("enter any number")
	fmt.Scanln(&num)
	for i := 1; i <= 10; i++ {
		table := num * i
		mul = mul * table
		fmt.Println(table)
	}
	fmt.Println("multiplication of oll digit is ", mul)
}
func NaturalNumber1To10() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}
func NaturalNumbersReverseOrder() {
	for i := 10; i >= 1; i-- {
		fmt.Println(i)
	}
}
func Divisibleby5and11() {
	var num int
	fmt.Println("enter any number")
	fmt.Scanln(&num)
	if num%5 == 0 || num%11 == 0 {
		fmt.Println(num, "is divisible")
	} else {
		fmt.Println(num, "is not divisible")
	}
}
func PrintOddNumbers1toN() {
	fmt.Println("printing odd number")
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}
func ProductOfNumber() {
	product := 1
	for i := 1; i < 10; i++ {
		product = product * i
		fmt.Println(i)
	}
	fmt.Println(product)
}
func PalindromeNumber() {
	var num, rem, temp int
	sum := 0
	fmt.Println("enter any number")
	fmt.Scanln(&num)
	temp = num
	for num > 0 {
		rem = num % 10
		sum = (sum * 10) + rem
		num = num / 10
	}
	if temp == sum {
		fmt.Println("is palindrome number")
	} else {
		fmt.Println("is not palindrome number")
	}
}
func PerfectNumber() {

}
func PrimeNumber() {
	var num int
	count := 0
	fmt.Println("enter any number")
	fmt.Scanln(&num)
	for i := 2; i < num; i++ {
		if num%i == 0 {
			count++
			break
		}
	}
	if count == 0 && num != 1 {
		fmt.Println(num, "is prime number")
	} else {
		fmt.Println(num, "is not prime number")
	}
}
func PositiveorNegative() {
	var num int
	fmt.Println("enter any number")
	fmt.Scanln(&num)
	if num > 0 {
		fmt.Println(num, "is positive number")
	} else {
		fmt.Println(num, "is negative number")
	}
}

func PowerOfANumber() {

	var powernum, exponentvalue float64

	fmt.Print("\nEnter the Number to find the Power = ")
	fmt.Scanln(&powernum)

	fmt.Print("Enter the Exponent Value = ")
	fmt.Scanln(&exponentvalue)

	power := math.Pow(powernum, exponentvalue)

	fmt.Println(powernum, " Power ", exponentvalue, " = ", power)
}
func ProfitOrLoss() {

	var profitcost, saleamount, amount int

	fmt.Print("\nEnter the Actual Product Cost = ")
	fmt.Scanln(&profitcost)

	fmt.Print("\nEnter the Sale Price = ")
	fmt.Scanln(&saleamount)

	if saleamount > profitcost {
		amount = saleamount - profitcost
		fmt.Println("Total Profit = ", amount)
	} else if profitcost > saleamount {
		amount = profitcost - saleamount
		fmt.Println("Total Loss = ", amount)
	} else {
		fmt.Println("No Profit No Loss")
	}
}
func Print1To100() {
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}
}
func ReverseNumber() {
	for i := 10; i >= 1; i-- {
		fmt.Println(i)
	}
}
func SimpleInterest() {
	var rate, time, principle float32
	fmt.Println("enter your amount")
	fmt.Scanln(&principle)
	fmt.Println("enter your rate")
	fmt.Scanln(&rate)
	fmt.Println("enter your time")
	fmt.Scanln(&time)
	si := (principle * rate * time) / 100
	fmt.Println(si)
}
func SquarerootofaNumber() {

	var sqrtnum, sqrtOut float64

	fmt.Print("\nEnter the Number to find Square Root = ")
	fmt.Scanln(&sqrtnum)

	sqrtOut = math.Sqrt(sqrtnum)

	fmt.Println("\nThe Square of a Given Number  = ", sqrtOut)
}
func SquareNum() {
	var num int
	fmt.Println("enter your number")
	fmt.Scanln(&num)
	square := num * num
	fmt.Println("square of num is := ", square)
}
func SumandAverageofNaturalNumbers() {
	var average, num int
	sum := 0
	count := 0
	fmt.Println("enter the maximum natural number")
	fmt.Scanln(&num)
	for i := 1; i <= num; i++ {
		sum = sum + i
		count++
	}
	average = sum / count
	fmt.Println("sum of total number is := ", sum)
	fmt.Println("average of total given number is := ", average)

}
func SumofDigitsinaNumber() {
	var reminder, num int
	sum := 0
	fmt.Println("enter two or more than digit")
	fmt.Scanln(&num)
	for num > 0 {
		num = num / 10
		sum = sum + reminder
		reminder = num % 10
	}
	fmt.Println("sum of all digit is := ", sum)
}
func SumOfEvenNumber() {
	sum := 0
	fmt.Println("printing all even number")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
			sum = sum + i
		}
	}
	fmt.Println("sum of even number is := ", sum)
}
func SumOfEvenAndOddNumber() {
	sum := 0
	sum1 := 0
	totalsum := 0
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("even number :", i)
			sum = sum + i
		} else {
			fmt.Println("odd number :", i)
			sum1 = sum1 + i
		}
		totalsum = sum + sum1
	}
	fmt.Println("sum of even and odd is :", totalsum)
}
func SumOfOddNumber() {
	sum := 0
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println(i)
			sum = sum + i
		}
	}
	fmt.Println("sum of all odd numbers is := ", sum)
}
func SwapTwoNumber() {
	var num1, num2, temp int
	fmt.Println("enter first number")
	fmt.Scanln(&num1)
	fmt.Println("enter Second number")
	fmt.Scanln(&num2)
	temp = num1
	num1 = num2
	num2 = temp
	fmt.Println("the first number after :", num1)
	fmt.Println("the second number after :", num2)
}
func SWAPTWONUMBERWITHOUTTEMPVARIABLE() {
	var a, b int
	fmt.Println("enter first number")
	fmt.Scanln(&a)
	fmt.Println("enter second number")
	fmt.Scanln(&b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Println("the first number after", a)
	fmt.Println("the second number after", b)
}
func swapTwonumberusingbitwise() {
	var a, b int
	fmt.Println("enter first number")
	fmt.Scanln(&a)
	fmt.Println("enter second number")
	fmt.Scanln(&b)
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println("the first number after", a)
	fmt.Println("the second number after", b)
}
