package main

import (
	"fmt"
)

/*
Add Two Arrays
Arithmetic Operations on an Array
Multiplication
Calculate Array Average
//Count Duplicates in an Array
Count Even and Odd Numbers in an Array
Count Positive and Negative Numbers in an Array
Largest Array Item
Largest and Smallest Number in an Array
Print Array Items in Even Index Position
Print Array Items in Odd Index Position
Print Even Numbers in an Array
Print Negative Numbers in an Array
Print Odd Numbers in an Array
Print Positive Numbers in an Array
Put Positive and Negatives in a Separate Array
Reverse an Array
Search for Array Items
Smallest Array Item
*/

func main() {
	//AddTwoArray()
	//ArithmeticOperations()
	//Multiplication()
	//CalculateAverageArray()
	//CountDublicate()
	//EvenAndOdd()
	//CountPositiveandNegative()
	//LargerstArray()
	//LargerstArray1() //user input
	//LargestandSmallest()
	//EvenIndexPosition()//dought hai
	//OddIndexPosition()//dought hai
	//EvenNumber()
	//OddNumber()
	//PositiveNumber()
	//PositiveOrNegative()
	//ReverseAnArray()
	//SearchForArrayItems()
	//SearchArry() //ye nahi bana
	//SmallestArrayItems()

}

func AddTwoArray() {
	var additionarray [5]int
	arr := [5]int{8, 7, 16, 5, 4}
	for i := 0; i < 5; i++ {
		fmt.Print(" ", arr[i])
	}
	fmt.Println()
	arr1 := [5]int{3, 13, 4, 5, 16}
	for i := 0; i < 5; i++ {
		fmt.Print(" ", arr1[i])
	}
	fmt.Println()
	for i, _ := range additionarray {
		additionarray[i] = arr[i] + arr1[i]
		fmt.Print(additionarray[i], "  ")
	}
}
func ArithmeticOperations() {
	arr1 := [5]int{12, 13, 14, 16, 18}
	arr2 := [5]int{8, 6, 7, 5, 4}
	fmt.Println("addition\tsubtraction\tmultiplication\tdivision\tmodule\treminder\t")
	for i := 0; i < 5; i++ {
		fmt.Print("\n", arr1[i]+arr2[i], "\t")
		fmt.Print("", arr1[i]-arr2[i], "\t")
		fmt.Print("", arr1[i]*arr2[i], "\t")
		fmt.Print("", arr1[i]%arr2[i], "\t")
		fmt.Print("", arr1[i]/arr2[i], "\t")
	}
}
func Multiplication() {
	var mularray [5]int
	var arr = [5]int{2, 3, 4, 5, 6}
	for i := 0; i < 5; i++ {
		fmt.Print("\t", arr[i])
	}
	fmt.Println()
	var arr1 = [5]int{9, 8, 7, 6, 5}
	for i := 0; i < 5; i++ {
		fmt.Print("\t ", arr1[i])
	}
	fmt.Println()
	for i := 0; i < len(mularray); i++ {
		mularray[i] = arr[i] * arr1[i]
	}
	fmt.Print(mularray, " ")
}
func CalculateAverageArray() {
	var avg int
	var sum = 0
	var count = 0
	arr := [5]int{10, 20, 30, 40, 50}
	for i := 0; i < 5; i++ {
		fmt.Println(arr[i])
		count++
		sum = sum + arr[i]
		avg = sum / count
	}
	fmt.Println("sum of array is ", sum)
	fmt.Println("average of array is ", avg)
}
func CountDublicate() {
	var dubsize, i int
	fmt.Println("enter the size of array")
	fmt.Scanln(&dubsize)

	duparr := make([]int, dubsize)

	fmt.Print("enter the even array item:")
	for i := 0; i < dubsize; i++ {
		fmt.Scan(&duparr[i])
	}
	dupcount := 0
	for i = 0; i < dubsize; i++ {
		for j := i + 1; j < dubsize; j++ {
			if duparr[i] == duparr[j] {
				dupcount++
				break
			}
		}
	}
	fmt.Println("\nThe Total Number of Duplicates in dupArr = ", dupcount)
}
func EvenAndOdd() {
	arr := [9]int{9, 5, 8, 7, 6, 3, 4, 2, 1}
	evencount := 0
	oddcount := 0
	for i := 0; i < 9; i++ {
		if arr[i]%2 == 0 {
			evencount++
			fmt.Printf("%d is even number", arr[i])
			fmt.Println()
		} else {
			fmt.Printf("%d is odd number", arr[i])
			oddcount++
			fmt.Println()
		}
	}
	fmt.Println("total even number is ", evencount)
	fmt.Println("total odd number is ", oddcount)
}
func CountPositiveandNegative() {
	positiveCount := 0
	negativeCount := 0
	arr := [9]int{2, -2, 3, 4, -5, -6, -7, -8, -9}
	for i := 0; i < 9; i++ {
		if arr[i] > 0 {
			positiveCount++
			fmt.Println("positive number", arr[i])
		} else {
			negativeCount++
			fmt.Println("negative number ", arr[i])
		}
	}
	fmt.Println("positive count is :=", positiveCount)
	fmt.Println("negative count is :=", negativeCount)
}
func LargerstArray() {
	var largerst, position int
	arr := [5]int{11, 17, 68, 43, 37}
	for i := 0; i < 5; i++ {
		fmt.Println(arr[i])
	}
	largerst = arr[0]
	for key, val := range arr {
		if largerst < val {
			largerst = val
			position = key
		}
	}
	fmt.Println("largest value is := ", largerst)
	fmt.Println("the index position of array is := ", position)
}
func LargerstArray1() {
	var largestsize, position int
	fmt.Println("enter the size of array")
	fmt.Scanln(&largestsize)

	arr := make([]int, largestsize)
	fmt.Println("enter the largest array items")
	for i, _ := range arr {
		fmt.Scanln(&arr[i])
	}
	largest := arr[0]
	for key, val := range arr {
		if largest < val {
			largest = val
			position = key
		}
	}
	fmt.Println("largest number of array is ", largest)
	fmt.Println("the index position of array is ", position)
}
func LargestandSmallest() {
	var largest, smallest, position, position1 int
	arr := [5]int{12, 44, 23, 75, 25}
	for i := 0; i < 5; i++ {
		fmt.Println(arr[i])
	}
	largest = arr[0]
	smallest = arr[0]
	for key, val := range arr {
		if largest < val {
			largest = val
			position = key
		}
		if smallest > val {
			smallest = val
			position1 = key
		}
	}
	fmt.Println("your largest value of array is ", largest)
	fmt.Println("the index value of array is ", position)
	fmt.Println()
	fmt.Println("your smallest value of array is ", smallest)
	fmt.Println("the index value of array is ", position1)
}
func EvenIndexPosition() {

	arr := [9]int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	fmt.Println("total even number")

	for key, _ := range arr {
		if key%2 == 0 {
			fmt.Print(" ", key)
			fmt.Println("\t", arr[key])

		}
	}

}
func OddIndexPosition() {
	arr := [5]int{9, 8, 7, 6, 5}
	for key, _ := range arr {
		if key%2 != 0 {
			fmt.Print(" key=", key)
			fmt.Println(" val=", arr[key])
		}
	}
}
func EvenNumber() {
	arr := [5]int{12, 13, 14, 15, 16}
	for i := 0; i < 5; i++ {
		if arr[i]%2 == 0 {
			fmt.Println("even number", arr[i])
		} else {
			fmt.Println("odd number", arr[i])
		}
	}
}
func OddNumber() {
	arr := [5]int{12, 13, 14, 15, 16}
	for i := 0; i < 5; i++ {
		if arr[i]%2 != 0 {
			fmt.Println("odd number is ", arr[i])
		}
	}
}
func PositiveNumber() {
	var size int
	fmt.Println("enter the size of array")
	fmt.Scanln(&size)

	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Println("enter element of array", i)
		fmt.Scanln(&arr[i])
	}
	fmt.Println()
	for i := 0; i < size; i++ {
		if arr[i] > 0 {
			fmt.Println("positive number", arr[i])
		}
	}
	fmt.Println()
}
func PositiveOrNegative() {
	posicount := 0
	negacount := 0
	arr := [5]int{6, 7, -3, -6, 2}
	for i := 0; i < 5; i++ {
		fmt.Print(arr[i])
	}
	fmt.Println("positive array")
	for i := 0; i < 5; i++ {
		if arr[i] >= 0 {
			posicount++
			fmt.Println(arr[i])
		}

	}
	fmt.Println("negative array")
	for i := 0; i < 5; i++ {
		if arr[i] <= 0 {
			negacount++
			fmt.Println(arr[i])
		}

	}
	fmt.Println("total number of positive array is ", posicount)
	fmt.Println("total number of negative array is ", negacount)
}
func ReverseAnArray() {
	var size, i int
	fmt.Println("enter the size of array")
	fmt.Scanln(&size)
	arr := make([]int, size)
	revarr := make([]int, size)

	fmt.Println("enter element of array")
	for i := 0; i < size; i++ {
		fmt.Scanln(&arr[i])

	}
	j := 0
	for i = size - 1; i >= 0; i-- {
		revarr[j] = arr[i]
		j++
	}
	fmt.Println("before array", arr)
	fmt.Println("reverse array is ", revarr)
}
func SearchForArrayItems() {
	var sersize, i, search int

	fmt.Print("Enter the Even Array Size = ")
	fmt.Scan(&sersize)

	serArr := make([]int, sersize)

	fmt.Print("Enter the Even Array Items  = \n")
	for i = 0; i < sersize; i++ {
		fmt.Scan(&serArr[i])
	}
	fmt.Print("Enter the Array Search Item  = ")
	fmt.Scan(&search)
	flag := 0
	for i = 0; i < sersize; i++ {
		if serArr[i] == search {
			flag = 1
			break
		}
	}
	if flag == 1 {
		fmt.Println("We Found the Search Item ", search, " at position = ", i)
	} else {
		fmt.Println("We haven't Found the Search Item ")
	}
}
func SearchArry() {
	var search int
	arr := [5]int{5, 9, 7, 8, 3}
	for i := 0; i < 5; i++ {
		fmt.Println(arr[i])
	}
	fmt.Println("enter your search item")
	fmt.Scanln(&search)

	flag := 0
	for i := 0; i < 5; i++ {
		if arr[i] == search {
			flag = 1
			break
		}
	}
	if flag == 1 {
		fmt.Println("we found the search item", search, "at position:=", i)
	} else {
		fmt.Println("we haven't found search item")
	}

}
func SmallestArrayItems() {
	var size, i, position int
	fmt.Println("enter the size of array")
	fmt.Scanln(&size)
	arr := make([]int, size)
	fmt.Println("enter the element of array")
	for i = 0; i < size; i++ {
		fmt.Scanln(&arr[i])

	}
	smallest := arr[0]
	for i := 0; i < size; i++ {
		if smallest > arr[i] {
			smallest = arr[i]
			position = i
		}
	}
	fmt.Println("smallest array items is ", smallest)
	fmt.Println("the index position of array is ", position)
}
