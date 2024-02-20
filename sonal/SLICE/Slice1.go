package main

import (
	"fmt"
	"sort"
)

/*
(1)Golang program to create a slice from an integer array
(2)Golang program to find the length of a slice
(3)Golang program to find the capacity of a slice
(4)Golang program to demonstrate the different ways to create slices
(5)Golang program to create a new slice from the existing slice
(6)Golang program to create a slice using the make() function
(7)Golang program to iterate a slice using 'range' in 'for' loop
(8)Golang program to iterate a slice using a range in 'for' loop without index
(9)Golang program to create and modify created slice
(10)Golang program to sort a slice of integer in ascending order
(11)Golang program to sort a slice of strings in ascending order
(12)Golang program to check a specified slice of integers is sorted or not
(13)Golang program to check a specified slice of strings is sorted or not
(14)Golang program to sort a slice of 64-bit floating-point numbers in ascending order
(15)Golang program to check a specified slice of 64-bit floating-point numbers is sorted or not
(16)Golang program to search an item in ascending order sorted slice
(17)Golang program to search an item in descending order sorted slice
(18)Golang program to search a floating-point number in a sorted slice using SearchFloat64s() function
(19)Golang program to search an integer item in a sorted slice using SearchInts() function
(20)Golang program to search a string in a sorted slice using SearchStrings() function
*/
func main() {
	//SliceInteger1()
	//DemonstrateTheDifferent()
	//NewSliceOrExisting()
	//SliceUsingMakeFunction()
	//IterateUsingLoop()
	//CreateAndModify()
	//SortSlice()
	//SortedSlice1()
	//CheckSortString()
	//SortSliFlotPoint()
	FloatSortOrNot()
}

func SliceInteger1() {
	Slice1 := []int{12, 13, 14, 16}
	for _, key := range Slice1 {
		fmt.Print(" ", key)
	}
	fmt.Println()
	fmt.Println("Len ", len(Slice1))
	fmt.Println("Cap ", cap(Slice1))
}

func DemonstrateTheDifferent() {
	Slice1 := []string{"Hello :", "my", "name", "is", "sonal", "kumar", "singh"}

	Slice2 := Slice1[1:]
	Slice3 := Slice1[0:5]
	Slice4 := Slice1[1:4]
	Slice5 := Slice1[:5]
	Slice6 := Slice1[5:]
	Slice7 := Slice1[1:3]

	fmt.Println("Slice2 :", Slice2)
	fmt.Println("Slice3 :", Slice3)
	fmt.Println("Slice4 :", Slice4)
	fmt.Println("Slice5 :", Slice5)
	fmt.Println("Slice6 :", Slice6)
	fmt.Println("Slice7 :", Slice7)
}

// Golang program to create a new slice
// from the existing slice

func NewSliceOrExisting() {
	//Create an array of integers.
	arr := [8]int{1, 2, 3, 4, 5, 6, 7, 8}

	OrgSlice := arr[1:7]
	NewSlice := OrgSlice[1:4]

	fmt.Println("Orginal slice: ", OrgSlice)
	fmt.Println("New slice: ", NewSlice)
}

func SliceUsingMakeFunction() {
	// Creating an array of size 8 and slice this array
	// till 5 and return the reference of the slice
	var slice = make([]int, 5, 8)

	fmt.Println("Slice: ", slice)
	fmt.Println("Length of Slice: ", len(slice))
	fmt.Println("Capacity of Slice: ", cap(slice))
}

func IterateUsingLoop() {
	var IntSli = []int{12, 14, 13, 17, 18, 16, 19}
	IntSli = IntSli[1:5]
	for index, value := range IntSli {
		fmt.Println("INDEX :", index, "VALUE :", value)
	}
	for _, value := range IntSli {
		fmt.Println("VALUE :", value)
	}
}

// Golang program to create and modify created slice

func CreateAndModify() {

	//Create an integer array
	arr := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	//create slice of from index 2 till index 4(5-1).
	intSlice := arr[2:5]

	intSlice[1] = 500
	fmt.Println("Slice elements: ")
	for _, ele := range intSlice {
		fmt.Printf("%d ", ele)
	}
}
func SortSlice() {
	SortedSlice := []int{12, 13, 15, 11, 9, 5, 6, 2}
	for i := 0; i < len(SortedSlice); i++ {
		fmt.Print(" ", SortedSlice[i])
	}
	fmt.Println()
	sort.Ints(SortedSlice)
	for _, ele := range SortedSlice {
		fmt.Print(" ", ele)
	}
}
func SortedSlice1() {
	SortString := []string{"Sonal", "Kumar", "Singh", "Harshit", "Chintu", "Asmirity"}
	for _, val := range SortString {
		fmt.Print(" ", val)

	}
	fmt.Println()
	sort.Strings(SortString)
	for i := 0; i < len(SortString); i++ {
		fmt.Print(" ", SortString[i])
	}
}
func CheckSortString() {
	var status bool
	Slice1 := []string{"Asmirity", "Chintu", "Harshit", "Sonal"}
	Slice2 := []string{"Sonal", "Kumar", "Singh"}
	status = sort.StringsAreSorted(Slice1)
	if status == true {
		fmt.Println("Slice1 is sorted")
	} else {
		fmt.Println("Slice1 is not sorted")
	}
	status = sort.StringsAreSorted(Slice2)
	if status == true {
		fmt.Println("Slice2 is sorted")
	} else {
		fmt.Println("Slice2 is not sorted")
	}
}
func SortSliFlotPoint() {
	Slice := []float64{86.6, 23.2, 32.2, 64.2, 68.4}
	sort.Float64s(Slice)
	for _, val := range Slice {
		fmt.Print(" ", val)
	}
}

func FloatSortOrNot() {
	var status bool
	var Slice1 = []float64{12.3, 13.4, 14.5, 15.6, 16.7, 17.8, 18.9}
	var Slice2 = []float64{12.3, 15.3, 82.6, 18.7, 58.9, 63.8, 32.7}
	status = sort.Float64sAreSorted(Slice1)
	if status == true {
		fmt.Print(" Slice1 is sorted")
	} else {
		fmt.Print(" Slice is not sorted")
	}
	fmt.Println()
	status = sort.Float64sAreSorted(Slice2)
	if status == true {
		fmt.Print("Slice2 is sorted")
	} else {
		fmt.Print(" Slice is not sorted")
	}

}
