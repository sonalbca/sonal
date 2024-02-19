package main

import (
	"fmt"
	"reflect"
)

func main() {
	ARRAY1()
	ARRAY2()
	ARRAY3()
	ARRAY4()
	ARRAY5()
	ARRAY6()
	ARRAY7()
	ARRAY8()
	ARRAY9()

}

/*----------------------------------- 1 -------------------------------------*/
/*
Golang Arrays
In this tutorial you'll learn how to store multiple values in a single variable in Golang

Introduction
An array is a data structure that consists of a collection of elements of a single type or simply you can say a special variable, which can hold more than one value at a time. The values an array holds are called its elements or items. An array holds a specific number of elements, and it cannot grow or shrink. Different data types can be handled as elements in arrays such as Int, String, Boolean, and others. The index of the first element of any dimension of an array is 0, the index of the second element of any array dimension is 1, and so on.
*/
//Declaring an Array

func ARRAY1() {
	var intArray [5]int
	var strArray [5]string

	fmt.Println(reflect.ValueOf(intArray).Kind())
	fmt.Println(reflect.ValueOf(strArray).Kind())
	/*
	   Output
	   array
	   array
	   array
	   array
	*/
}

/*-------------------------------------- 2 ------------------------------------*/
//Assign and Access Values
//You access or assign the array elements by referring to the index number. The index is specified in square brackets.

func ARRAY2() {
	var theArray [3]string
	theArray[0] = "India"  // Assign a value to the first element
	theArray[1] = "Canada" // Assign a value to the second element
	theArray[2] = "Japan"  // Assign a value to the third element

	fmt.Println(theArray[0]) // Access the first element value
	fmt.Println(theArray[1]) // Access the second element valu
	fmt.Println(theArray[2]) // Access the third element valu
	/*
	   Output
	   India
	   Canada
	   Japan
	*/
}

/*------------------------------------------ 3 ---------------------------------*/
//Initializing an Array with an Array Literal
//You can initialize an array with pre-defined values using an array literal. An array literal have the number of elements it will hold in square brackets, followed by the type of its elements. This is followed by a list of initial values separated by commas of each element inside the curly braces.

func ARRAY3() {
	x := [5]int{10, 20, 30, 40, 50}   // Intialized with values
	var y [5]int = [5]int{10, 20, 30} // Partial assignment

	fmt.Println(x)
	fmt.Println(y)
	/*
	   Output
	   [10 20 30 40 50]
	   [10 20 30 0 0]
	*/
}

/*------------------------------------------- 4 --------------------------------*/
//Initializing an Array with ellipses...
//When we use ... instead of specifying the length. The compiler can identify the length of an array, based on the elements specified in the array declaration.

func ARRAY4() {
	x := [...]int{10, 20, 30}

	fmt.Println(reflect.ValueOf(x).Kind())
	fmt.Println(len(x))
	/*
	   Output
	   array
	   3
	*/
}

/*----------------------------------------- 5 -----------------------------------*/
//Initializing Values for Specific Elements
//When an array declare using an array literal, values can be initialize for specific elements.

//A value of 10 is assigned to the second element (index 1) and a value of 30 is assigned to the fourth element (index 3).

func ARRAY5() {
	x := [5]int{1: 10, 3: 30}
	fmt.Println(x)
	/*
	   Output
	   [0 10 0 30 0]
	*/
}

/*------------------------------------------ 6 -----------------------------------*/
//Loop Through an Indexed Array
//You can loop through an array elements by using a for loop.

func ARRAY6() {
	intArray := [5]int{10, 20, 30, 40, 50}

	fmt.Println("\n---------------Example 1--------------------\n")
	for i := 0; i < len(intArray); i++ {
		fmt.Println(intArray[i])
	}

	fmt.Println("\n---------------Example 2--------------------\n")
	for index, element := range intArray {
		fmt.Println(index, "=>", element)

	}

	fmt.Println("\n---------------Example 3--------------------\n")
	for _, value := range intArray {
		fmt.Println(value)
	}

	j := 0
	fmt.Println("\n---------------Example 4--------------------\n")
	for range intArray {
		fmt.Println(intArray[j])
		j++
	}
}

/*--------------------------------------- 7 ----------------------------------*/
/*
Copy Array
You can create copy of an array, by assigning an array to a new variable either by value or reference.
*/

func ARRAY7() {

	strArray1 := [3]string{"Japan", "Australia", "Germany"}
	strArray2 := strArray1  // data is passed by value
	strArray3 := &strArray1 // data is passed by refrence

	fmt.Printf("strArray1: %v\n", strArray1)
	fmt.Printf("strArray2: %v\n", strArray2)

	strArray1[0] = "Canada"

	fmt.Printf("strArray1: %v\n", strArray1)
	fmt.Printf("strArray2: %v\n", strArray2)
	fmt.Printf("*strArray3: %v\n", *strArray3)
	/*
		Output
		strArray1: [Japan Australia Germany]
		strArray2: [Japan Australia Germany]
		strArray1: [Canada Australia Germany]
		strArray2: [Japan Australia Germany]
		*strArray3: [Canada Australia Germany]
	*/
}

/*----------------------------------------- 8 -------------------------------------*/
//Check if Element Exists

/*This Go program checks whether an element exists in an array or not.
The program takes an array and an item as input arguments,
and it returns a boolean value indicating whether the item exists in the array or not.*/

/*The program uses the reflect package in Go to determine the kind of data type of the input array.
If the input is not an array, the program panics with an "Invalid data-type" error.*/

/*The itemExists function in the program uses a loop to iterate through the elements of the input array.
It uses the Index method of the reflect.Value type to access the value of the array element at each index.
It then compares the value with the input item using the Interface method of reflect.Value.
If a match is found, the function returns true, indicating that the item exists in the array.
If no match is found, the function returns false.*/

/*In the main function, the program creates an array of strings and calls the itemExists function twice to
check whether the items "Canada" and "Africa" exist in the array.
The program prints the return values of the function calls,
which are either true or false, depending on whether the items exist in the array or not.*/

func ARRAY8() {
	strArray := [5]string{"India", "Canada", "Japan", "Germany", "Italy"}
	fmt.Println(itemExists(strArray, "Canada"))
	fmt.Println(itemExists(strArray, "Africa"))
}

func itemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	if arr.Kind() != reflect.Array {
		panic("Invalid data-type")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false

	/*
	   Output
	   true
	   false
	*/
}

/*-------------------------------------- 9 -------------------------------------*/
/*
Filter Array Elements
You can filter array element using : as shown below

A value of 10 is assigned to the second element (index 1) and a value of 30 is assigned to the fourth element (index 3).
*/

func ARRAY9() {
	countries := [...]string{"India", "Canada", "Japan", "Germany", "Italy"}

	fmt.Printf("Countries: %v\n", countries)

	fmt.Printf(":2 %v\n", countries[:2])

	fmt.Printf("1:3 %v\n", countries[1:3])

	fmt.Printf("2: %v\n", countries[2:])

	fmt.Printf("2:5 %v\n", countries[2:5])

	fmt.Printf("0:3 %v\n", countries[0:3])

	fmt.Printf("Last element: %v\n", countries[len(countries)-1])

	fmt.Printf("All elements: %v\n", countries[0:len(countries)])
	fmt.Println(countries[:])
	fmt.Println(countries[0:])
	fmt.Println(countries[0:len(countries)])

	fmt.Printf("Last two elements: %v\n", countries[len(countries)-2:len(countries)])
	/*
	   Output
	   Countries: [India Canada Japan Germany Italy]
	   :2 [India Canada]
	   1:3 [Canada Japan]
	   2: [Japan Germany Italy]
	   2:5 [Japan Germany Italy]
	   0:3 [India Canada Japan]
	   Last element: Italy
	   All elements: [India Canada Japan Germany Italy]
	   [India Canada Japan Germany Italy]
	   [India Canada Japan Germany Italy]
	   [India Canada Japan Germany Italy]
	   Last two elements: [Germany Italy]
	*/
}
