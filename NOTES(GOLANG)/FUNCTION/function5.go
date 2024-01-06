package main

import "fmt"

/*
Naming Conventions for Golang Functions

A name must begin with a letter, and can have any number of additional letters and numbers.
A function name cannot start with a number.
A function name cannot contain spaces.
If the functions with names that start with an uppercase letter will be exported to other packages. If the function name starts with a lowercase letter, it won't be exported to other packages, but you can call this function within the same package.
If a name consists of multiple words, each word after the first should be capitalized like this: empName, EmpAddress, etc.
function names are case-sensitive (car, Car and CAR are three different variables).
*/

//Golang Passing Address to a Function

//Passing the address of variable to the function and the value of a variables modified using dereferencing inside body of function.

func update(a *int, t *string) {
	*a = *a + 5      // defrencing pointer address
	*t = *t + " Doe" // defrencing pointer address
	return
}

func main() {
	var age = 20
	var text = "John"
	fmt.Println("Before:", text, age)

	update(&age, &text)

	fmt.Println("After :", text, age)
}

/*
Output
Before: John 20
After : John Doe 25
*/
