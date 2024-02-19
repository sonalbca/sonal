package main

import (
	"fmt"
	"reflect"
)

func main() {
	//Variable1()
	//Variable2()
	//Variable3()
	//Variable4()
	//Variable5()
	Variable6()
	//Variable7()
	//Variable8()
}

/*
Golang Variables
Variables
A variable is a piece of storage containing data temporarily to work with it.

The most general form to declare a variable in Golang uses the var keyword, an explicit type, and an assignment.

var name type = expression
*/

//Variable Declaration with Initialization
//If you know beforehand what a variable's value will be, you can declare variables and assign them values on the same line.

func Variable1() {
	var i = 10
	var s = "Canada"

	fmt.Println(i)
	fmt.Println(s)
}

/*
Inside the main function:
var i int = 10: This declares a variable named i of type int and initializes it with the value 10.
var s string = "Canada": This declares a variable named s of type string and initializes it with the value "Canada".

The Go language has strong static typing, which means you need to specify the type of a variable when you declare it.
*/

//Variable Declaration without Initialization
//The keyword var is used for declaring variables followed by the desired name and the type of value the variable will hold.

func Variable2() {
	var i int
	var s string

	i = 10
	s = "Canada"

	fmt.Println(i)
	fmt.Println(s)
}

/*
Two variables are declared here:
- i of type int
- s of type string
At this point, they're declared but not initialized with meaningful values.
The variables declared earlier are now assigned values:
- The integer 10 is assigned to i.
- The string "Canada" is assigned to s.
*/

//Variable Declaration with type inference
//You can omit the variable type from the declaration, when you are assigning a value to a variable at the time of declaration. The type of the value assigned to the variable will be used as the type of that variable.

func Variable3() {
	var i = 10
	var s = "Canada"

	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(s))
}

/*
Here, two variables are declared and immediately initialized:
- i is initialized with the integer value 10.
- s is initialized with the string value "Canada".
Note: Since Go supports type inference, you don't have to explicitly mention the type while initializing the variable. Go automatically determines the type based on the value.
Here, the TypeOf function from the reflect package is used to determine the type of the variables i and s at runtime.
- reflect.TypeOf(i) returns the type of variable i, which is int.
- reflect.TypeOf(s) returns the type of variable s, which is string.
These types are then printed to the console using the fmt.Println function.
*/

//Multiple Variable Declaration
//Golang allows you to assign values to multiple variables in one line.

func Variable4() {
	var fname, lname = "John", "Doe"
	m, n, o := 1, 2, 3
	item, price := "Mobile", 2000

	fmt.Println(fname + lname)
	fmt.Println(m + n + o)
	fmt.Println(item, "-", price)
}

/*
Two string variables, fname and lname, are declared and initialized simultaneously. fname is assigned the value "John" and lname is assigned the value "Doe".

Here, three integer variables (m, n, and o) are declared and initialized using Go's short variable declaration syntax (:=). The values are 1, 2, and 3 respectively.

Similarly, two variables, item (a string) and price (an integer), are declared and initialized. item is assigned the value "Mobile" and price is assigned the value 2000.
*/

//Short Variable Declaration
//The := short variable assignment operator indicates that short variable declaration is being used. There is no need to use the var keyword or declare the variable type.

func Variable5() {
	name := "John Doe"
	fmt.Println(reflect.TypeOf(name))
}

/*
Here, a string variable named name is declared and initialized with the value "John Doe" using Go's short variable declaration syntax (:=).
*/

//Scope of Golang Variables Defined by Brace Brackets
//In the Go language, the scope of a variable refers to the region of the code where the variable is accessible. The scope of a variable is determined by where it is declared.

//A variable declared within a block (e.g., inside a function or a control structure such as if, for, or switch) is only accessible within that block and any nested blocks. Once the execution leaves the block, the variable goes out of scope and cannot be accessed anymore.

var s = "Japan"

func Variable6() {
	fmt.Println(s) // Accessing the global variable 's' and printing its value
	x := true      // Declaring and initializing a local variable 'x'

	if x { // If 'x' is true, execute the block of code inside the if statement
		y := 1 // Declaring and initializing a local variable 'y' inside the if block

		if x != false { // If 'x' is not false, execute the block of code inside the if statement
			fmt.Println(s) // Accessing the global variable 's' and printing its value
			fmt.Println(x) // Printing the value of local variable 'x'
			fmt.Println(y) // Printing the value of local variable 'y'
		}
	}

	fmt.Println(x) // Printing the value of local variable 'x'
}

/*
In the main function, we have three variables with different scopes:
s: This is a global variable and can be accessed from any function within the package, including the main function.
x: This is a local variable declared and initialized inside the main function. It is accessible only within the main function. In this case, it is also accessible inside the if blocks within the main function.
y: This is a local variable declared and initialized inside the first if block. Its scope is limited to the block it is declared in, so it is accessible only within that if block and its nested blocks. In this case, it is also accessible inside the second if block, which is nested inside the first if block.
*/

/*
Golang Variables Naming Conventions
These are the following rules for naming a Golang variable:

A name must begin with a letter, and can have any number of additional letters and numbers.
A variable name cannot start with a number.
A variable name cannot contain spaces.
If the name of a variable begins with a lower-case letter, it can only be accessed within the current package this is considered as unexported variables.
If the name of a variable begins with a capital letter, it can be accessed from packages outside the current package one this is considered as exported variables.
If a name consists of multiple words, each word after the first should be capitalized like this: empName, EmpAddress, etc.
Variable names are case-sensitive (car, Car and CAR are three different variables).
*/
//Zero value in Golang
//If you declare a variable without assigning it a value, Golang will automatically bind a default value (or a zero-value) to the variable.

func Variable7() {
	var quantity float32
	fmt.Println(quantity)

	var price int16
	fmt.Println(price)

	var product string
	fmt.Println(product)

	var inStock bool
	fmt.Println(inStock)
}

/*
The program is written in the Go language and declares four variables of different data types: quantity (float32), price (int16), product (string), and inStock (bool).
Each variable is then initialized with its zero value according to its data type.
After the initialization, the program prints the value of each variable to the console in the order they are declared. The output of the program will be:

0
0

false
*/

//Golang Variable Declaration Block
//Variables declaration can be grouped together into blocks for greater readability and code quality.

var (
	product  = "Mobile"
	quantity = 50
	price    = 50.50
	inStock  = true
)

func Variable8() {
	fmt.Println(quantity)
	fmt.Println(price)
	fmt.Println(product)
	fmt.Println(inStock)
}

/*
Global Variables Declaration: A group of global variables is declared using the var keyword with a block of parentheses. These variables are accessible throughout the program, not just within the main function.:
product (string): It holds the product name, which is "Mobile".
quantity (int): It holds the number of products, which is 50.
price (float64): It holds the price of the product, which is 50.50.
inStock (bool): It holds a boolean value indicating the stock status, which is true (i.e., the product is in stock).
*/
