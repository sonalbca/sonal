package main

import "fmt"

var positiveCount, negativeCount, i int

func printArray(posNegArr []int, size int) {
	for i = 0; i < size; i++ {
		fmt.Print(posNegArr[i], " ")
	}
}
func putPositiveNums(posNegArr []int, size int) {
	posArr := make([]int, size)
	positiveCount = 0
	for i = 0; i < size; i++ {
		if posNegArr[i] >= 0 {
			posArr[positiveCount] = posNegArr[i]
			positiveCount++
		}
	}
	fmt.Println("The Total Number of Positive Numbers = ", positiveCount)
	fmt.Print("The Positive Array Elements are = ")
	printArray(posArr, positiveCount)
}
func putNegativeNums(posNegArr []int, size int) {
	NegArr := make([]int, size)
	negativeCount = 0
	for i = 0; i < size; i++ {
		if posNegArr[i] < 0 {
			NegArr[negativeCount] = posNegArr[i]
			negativeCount++
		}
	}
	fmt.Println("\nThe Total Number of Negative Numbers = ", negativeCount)
	fmt.Print("The Negative Array Elements are = ")
	printArray(NegArr, negativeCount)
}
func main() {
	var size, i int

	fmt.Print("Enter the Positive Negative Array Size = ")
	fmt.Scan(&size)

	posNegArr := make([]int, size)

	fmt.Print("Enter the Positive Negative Array Items  = ")
	for i = 0; i < size; i++ {
		fmt.Scan(&posNegArr[i])
	}

	putPositiveNums(posNegArr, size)
	putNegativeNums(posNegArr, size)
	fmt.Println()

}
