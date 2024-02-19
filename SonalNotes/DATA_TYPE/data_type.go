package main

import (
	"bytes"
	"fmt"
	"strings"
)

/*
Data Types
Go is a statically typed programming language. This means that variables always have a specific type and that type cannot change.
The keyword var is used for declaring variables of a particular data type. Here is the syntax for declaring variables:

var name type = expression
On the left we use the var keyword to declare a variable and then assign a value to it.
We can declare mutiple variables of the same type in a single statement as shown here:

var fname,lname string
Multiple variables of the same type can also be declared on a single line: var x, y int makes x and y both int variables.
You can also make use of parallel assignment: a, b := 20, 16 If you are using an initializer expression for declaring variables,
you can omit the type using short variable declaration as shown here:

country, state := "Germany", "Berlin"
We use the operator : = for declaring and initializing variables with short variable declaration.
When you declare variables with this method, you can't specify the type because the type is determined by the initializer expression.
*/

func main() {
	Data1()

	//Data2()
	//Data3()
	//Data4()
	//Data6()
	//Data7()

	//DATA8()

}

// ------------------------------ 1 ------------------------------//
// Global variable declaration
var (
	m int
	n int
)

func Data1() {
	var x int = 1 // Integer Data Type
	var y int     //  Integer Data Type
	fmt.Println(x)
	fmt.Println(y)

	var a, b, c = 5.25, 25.25, 14.15 // Multiple float32 variable declaration
	fmt.Println(a, b, c)

	city := "Berlin"     // String variable declaration
	Country := "Germany" // Variable names are case sensitive
	fmt.Println(city)
	fmt.Println(Country) // Variable names are case sensitive

	food, drink, price := "Pizza", "Pepsi", 125 // Multiple type of variable declaration in same line
	fmt.Println(food, drink, price)
	m, n = 1, 2
	fmt.Println(m, n)
}

//------------------------------ 2 ------------------------------//
/*
Golang String Concatenation
Strings are a fundamental building block of programming. String concatenation means add strings together.

String Concatenation Using + (Concatenation Operator)
// Golang program to concatenate strings
*/

func Data2() {
	// Creating and initializing strings
	// using var keyword
	var str1 string
	str1 = "Great"

	var str2 string
	str2 = "Britain"

	// Concatenating strings
	// Using + operator
	fmt.Println(str1 + str2)

	// Creating and initializing strings
	// Using shorthand declaration
	str3 := "New"
	str4 := "York"

	// Concatenating strings
	// Using + operator
	result := str3 + " " + str4

	fmt.Println(result)
}

/*
This Golang program demonstrates how to concatenate strings. It first initializes two strings str1 with the value "Great" and str2 with the value "Britain" using the var keyword.

These strings are then concatenated using the + operator and the result "GreatBritain" is printed to the console.

The program then uses a shorthand declaration to initialize another two strings str3 with the value "New" and str4 with the value "York". These strings are concatenated with a space in between, resulting in "New York", which is then printed to the console.

Append String Using += operator

*/
// Golang program to illustrate
// how to append string

func Data3() {

	// Creating and initializing strings
	str1 := "New"
	str2 := "York"

	// Using += operator
	str1 += str2
	fmt.Println(str1)

	str1 += "London"
	fmt.Println(str1)
}

/*
The Golang program demonstrates how to append strings. It initializes two strings, str1 with the value "New" and str2 with the value "York".

The program then appends str2 to str1 using the += operator and prints the result, which is "NewYork".

After that, the program appends the string "London" to the modified str1 and prints the resulting string, which is "NewYorkLondon".

String Concatenation Using Strings Builder
*/

// Golang program to illustrate
// how to concatenate strings
// Using strings.Builder with WriteString() function

func Data4() {
	// Creating and initializing strings
	// Using strings.Builder with
	// WriteString() function
	var sb strings.Builder

	sb.WriteString("London")
	sb.WriteString("-")
	sb.WriteString("5214")

	fmt.Println(sb.String())
}

/*
This Golang program demonstrates string concatenation using the strings.Builder type and its WriteString() method.

The program initializes a strings.Builder variable named sb and then appends the string "London", followed by a dash "-", and then the string "5214" to it.

The concatenated result "London-5214" is then printed to the standard output using the fmt.Println() function.

String Concatenation Using Bytes Buffer
*/

// Golang program to illustrate
// how to concatenate strings
// Using bytes.Buffer with WriteString() function

func Data5() {
	// Creating and initializing strings
	// Using bytes.Buffer with
	// WriteString() function
	var b bytes.Buffer

	b.WriteString("London")
	b.WriteString("#")
	b.WriteString("5214")

	fmt.Println(b.String())
}

/*
The given Golang program demonstrates how to concatenate strings using the bytes.Buffer type and its WriteString() method.

The program uses the bytes.Buffer type to concatenate three strings: "London", "#", and "5214".

It initializes an empty buffer b and then appends each of these strings to the buffer using the WriteString() method.

Finally, the concatenated result "London#5214" is printed to the console using the fmt.Println() function.

String Concatenation Using Sprintf Method
*/

// Golang program to illustrate
// how to concatenate strings
// Using Sprintf() method

func Data6() {
	// Creating and initializing strings
	s1 := "London"
	s2 := "@"
	s3 := "5421"

	result := fmt.Sprintf("%s%s%s", s1, s2, s3)

	fmt.Println(result)
}

/*
This Golang program demonstrates how to concatenate multiple strings using the Sprintf() method from the "fmt" package.

The program initializes three strings: s1 with the value "London", s2 with the value "@", and s3 with the value "5421". It then concatenates these strings using fmt.Sprintf() and stores the result in the variable result.

Finally, the program prints the concatenated result, which will be "London@5421".

Repeat same String Multiple Times
*/

// Golang program to illustrate
// how to repeat same string multiple times
// using Sprintf() method

func Data7() {
	// Creating and initializing strings
	s1 := "London"
	r := fmt.Sprintf("%s", strings.Repeat(s1, 3))
	fmt.Println(r)
}

/*
This Golang program demonstrates how to repeat a string multiple times using the Sprintf() method from the fmt package and the Repeat() function from the strings package.

The program defines a string s1 with the value "London". It then repeats the string three times using the strings.Repeat() function. The repeated string is then formatted using fmt.Sprintf() and assigned to the variable r.

Finally, the program prints the repeated string (which would be "LondonLondonLondon") to the console.
*/

//----------------------------- RUNE -------------//

//What is Rune? How to get ASCII value of any character in Go?

func DATA8() {
	// Example - 1
	str := "GOLANG"
	runes := []rune(str)

	var result []int

	for i := 0; i < len(runes); i++ {
		result = append(result, int(runes[i]))
	}

	fmt.Println(result)

	// Example - 2
	s := "GOLANG"
	for _, r := range s {
		fmt.Printf("%c - %d\n", r, r)
	}
}
