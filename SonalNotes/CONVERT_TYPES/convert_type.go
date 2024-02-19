package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

/*
Type Conversion
The process of converting a value from one data type to another is known as type conversion.
Converting from one datatype to another is a frequent task we programmers do.
Let's take few examples of converting variables from one data type to another.
*/

func main() {
	//Convert1()
	//Convert2()
	//Convert3()
	//Convert4()
	//Convert5()
	//Convert6()
	//Convert7()
	//Convert8()
	//Convert9()
	//Convert10()
	//Convert11()
	Convert12()
	Convert13()
	Convert14()
	Convert15()

}

//------------------------------- 1 -----------------------------//
//How to Convert string to integer type in Go?
/*
Like most modern languages, Golang includes strings as a built-in type. Let's take an example,
you may have a string that contains a numeric value "100". However, because this value is represented as a string,
you can't perform any mathematical calculations on it. You need to explicitly convert this string type into an integer
type before you can perform any mathematical calculations on it. In order to convert string to integer type in Golang,
you can use the following methods.

Atoi() Function

You can use the strconv package's Atoi() function to convert the string into an integer value.
Atoi stands for ASCII to integer. The Atoi() function returns two values: the result of the conversion, and the error (if any).
*/

//syntax
//func Atoi(s string) (int, error)

func Convert1() {
	strVar := "100"
	intVar, err := strconv.Atoi(strVar)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))

	/*Output
	100 < nil > int
	*/
}

/*
ParseInt() Function
ParseInt interprets a string s in the given base (0, 2 to 36) and bit size (0 to 64) and returns the corresponding value i. This function accepts a string parameter, convert it into a corresponding int type based on a base parameter. By default, it returns Int64 value.

Syntax
func ParseInt(s string, base int, bitSize int) (i int64, err error)
*/

func Convert2() {
	strVar := "100"

	intVar, err := strconv.ParseInt(strVar, 0, 8)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))

	intVar, err = strconv.ParseInt(strVar, 0, 16)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))

	intVar, err = strconv.ParseInt(strVar, 0, 32)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))

	intVar, err = strconv.ParseInt(strVar, 0, 64)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))
	/*
		Output
		100 < nil > int64
		100 < nil > int64
		100 < nil > int64
		100 < nil > int64
	*/
}

//Using fmt.Sscan
//The fmt package provides sscan() function which scans string argument and store into variables. This function read the string with spaces and assign into consecutive Integer variables.

func Convert3() {
	strVar := "100"
	intValue := 0
	_, err := fmt.Sscan(strVar, &intValue)
	fmt.Println(intValue, err, reflect.TypeOf(intValue))

	/*
		Output
		100 < nil > int
	*/
}

//------------------------------- 2 -----------------------------//
/*
How to Convert string to float type in Go?

Like most modern languages, Golang includes strings as a built-in type.
Let's take an example, you may have a string that contains a numeric value "3.1415926535".
However, because this value is represented as a string, you can't perform any mathematical calculations on it.
You need to explicitly convert this string type into an float type before you can perform any mathematical calculations on it.

ParseFloat() Function
ParseFloat converts the string s to a floating-point number with the precision specified by bitSize: 32 for float32, or 64 for float64.
When bitSize=32, the result still has type float64, but it will be convertible to float32 without changing its value.

Syntax
func ParseFloat(s string, bitSize int) (float64, error)
*/

func Convert4() {
	s := "3.1415926535"
	f, err := strconv.ParseFloat(s, 8)
	fmt.Println(f, err, reflect.TypeOf(f))

	s1 := "-3.141"
	f1, err := strconv.ParseFloat(s1, 8)
	fmt.Println(f1, err, reflect.TypeOf(f1))

	s2 := "A-3141X"
	f2, err := strconv.ParseFloat(s2, 32)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f2, err, reflect.TypeOf(f2))
	}
}

//To check if an error occurred during the conversion, check if the err variable contains a nil value.

/*
Output
3.1415926535 <nil> float64
-3.141 <nil> float64
strconv.ParseFloat: parsing "A-3141X": invalid syntax
*/

//------------------------------- 3 -----------------------------//
/*
How to convert String to Boolean Data Type Conversion in Go?

Like most modern languages, Golang includes strings as a built-in type.
Let's take an example, you may have a string that contains a boolean value "true".
However, because this value is represented as a string, you can't perform any operation on it.
You need to explicitly convert this string type into an boolean type before you can perform any operation on it.


String to Boolean Conversion

Package strconv is imported to perform conversions to and from string.ParseBool returns the boolean value represented by the string.
It accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False. Any other value returns an error.
*/

func Convert5() {
	s1 := "true"
	b1, _ := strconv.ParseBool(s1)
	fmt.Printf("%T, %v\n", b1, b1)

	s2 := "t"
	b2, _ := strconv.ParseBool(s2)
	fmt.Printf("%T, %v\n", b2, b2)

	s3 := "0"
	b3, _ := strconv.ParseBool(s3)
	fmt.Printf("%T, %v\n", b3, b3)

	s4 := "F"
	b4, _ := strconv.ParseBool(s4)
	fmt.Printf("%T, %v\n", b4, b4)
}

/*
Output
bool, true
bool, true
bool, false
bool, false
*/

//------------------------------- 4 -----------------------------//
/*
How to convert Boolean Type to String in Go?

Like most modern languages, Golang includes Boolean as a built-in type. Let's take an example, you may have a variable that contains a boolean value true. In order to convert boolean vale into string type in Golang, you can use the following methods.

FormatBool function

You can use the strconv package's FormatBool() function to convert the boolean into an string value. FormatBool returns "true" or "false" according to the value of b.
Syntax
func FormatBool(b bool) string
*/

func Convert6() {
	var b bool = true
	fmt.Println(reflect.TypeOf(b))

	var s string = strconv.FormatBool(true)
	fmt.Println(reflect.TypeOf(s))
}

/*
Output
bool
string

*/

/*fmt.Sprintf() method
Sprintf formats according to a format specifier and returns the resulting string. Here, a is of Interface type hence you can use this method to convert any type to string.
Syntax
func Sprintf(format string, a ...interface{}) string
*/

func Convert7() {
	b := true
	s := fmt.Sprintf("%v", b)
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))
}

/*
Output
true
string
*/

//------------------------------- 5 -----------------------------//
/*
How to Convert Float to String type in Go?
Like most modern languages, Golang includes Float as a built-in type.
Let's take an example, you may have a variable that contains a Float value.
In order to convert Float value into string type in Golang, you can use the following methods.


FormatFloat method
You can use the strconv package's FormatFloat() function to convert the float into an string value.
FormatFloat converts the floating-point number f to a string, according to the format fmt and precision prec.
It rounds the result assuming that the original was obtained from a
floating-point value of bitSize bits (32 for float32, 64 for float64).

Syntax
func FormatFloat(f float64, fmt byte, prec, bitSize int) string
*/

func Convert8() {
	var f float64 = 3.1415926535
	fmt.Println(reflect.TypeOf(f))
	fmt.Println(f)

	var s string = strconv.FormatFloat(f, 'E', -1, 32)
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(s)
}

/*
Output
float64
3.1415926535
string
3.1415927E+00
*/
//------------------------------------------------------------------------//
/*
fmt.Sprintf() method
Sprintf formats according to a format specifier and returns the resulting string.
Here, a is of Interface type hence you can use this method to convert any type to string.

Syntax
func Sprintf(format string, a ...interface{}) string
*/

func Convert9() {
	b := 12.454
	fmt.Println(reflect.TypeOf(b))

	s := fmt.Sprintf("%v", b)
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))
}

/*
Output
float64
12.454
string
*/

//------------------------------- 6 -----------------------------//
/*
Different ways for Integer to String Conversions
Like most modern languages, Golang includes Integer as a built-in type.
Let's take an example, you may have a variable that contains a Integer value and you want to convert it into String.
In order to convert Integer value into String type in Golang, you can use the following methods.

FormatInt() Method
You can use the strconv package's FormatInt() function to convert the int into an string value.
FormatInt returns the string representation of i in the given base, for 2 <= base <= 36.
The result uses the lower-case letters 'a' to 'z' for digit values >= 10.

Syntax
func FormatInt(i int64, base int) string
*/

func Convert10() {
	var i int64 = 125
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(i)

	var s string = strconv.FormatInt(i, 10)
	fmt.Println(reflect.TypeOf(s))

	fmt.Println("Base 10 value of s:", s)

	s = strconv.FormatInt(i, 8)
	fmt.Println("Base 8 value of s:", s)

	s = strconv.FormatInt(i, 16)
	fmt.Println("Base 16 value of s:", s)

	s = strconv.FormatInt(i, 32)
	fmt.Println("Base 32 value of s:", s)
}

/*
Output
int64
125
string
Base 10 value of s: 125
Base 8 value of s: 175
Base 16 value of s: 7d
Base 32 value of s: 3t
*/

//-----------------------------------------------------------//
/*
fmt.Sprintf() Method
Sprintf formats according to a format specifier and returns the resulting string.
Here, a is of Interface type hence you can use this method to convert any type to string.

	Syntax
 func Sprintf(format string, a ...interface{}) string
*/

func Convert11() {
	b := 1225
	fmt.Println(reflect.TypeOf(b))

	s := fmt.Sprintf("%v", b)
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))
}

/*
Output
int
1225
string
*/

//------------------------------- 7 -----------------------------//

//Convert Int data type to Int16 Int32 Int64

func Convert12() {
	var i int = 10
	fmt.Println(reflect.TypeOf(i))

	i16 := int16(i)
	fmt.Println(reflect.TypeOf(i16))

	i32 := int32(i)
	fmt.Println(reflect.TypeOf(i32))

	i64 := int64(i)
	fmt.Println(reflect.TypeOf(i64))
}

/*
Output
int
int16
int32
int64
*/

//------------------------------- 7 -----------------------------//
//Convert Float32 to Float64 and Float64 to Float32

func Convert13() {
	var f32 float32 = 10.6556
	fmt.Println(reflect.TypeOf(f32))

	f64 := float64(f32)
	fmt.Println(reflect.TypeOf(f64))

	f64 = 1097.655698798798
	fmt.Println(f64)

	f32 = float32(f64)
	fmt.Println(f32)
}

/*
Output
float32
float64
1097.655698798798
1097.6556
*/

//------------------------------- 8 -----------------------------//
//Converting Int data type to Float in Go

func Convert14() {
	var f32 float32 = 10.6556
	fmt.Println(reflect.TypeOf(f32))

	i32 := int32(f32)
	fmt.Println(reflect.TypeOf(i32))
	fmt.Println(i32)

	f64 := float64(i32)
	fmt.Println(reflect.TypeOf(f64))
}

/*
Output
float32
int32
10
float64
*/

//------------------------------- 8 -----------------------------//
//------------------------------- 8 -----------------------------//
/*
How to trim leading and trailing white spaces of a string in Golang?
Removing leading and trailing spaces from a string eliminates all whitespace before the first character and after the last character in the string.
For example, removing leading and trailing spaces from " abc def " results in "abc def". The strings.TrimSpace() method from the standard library, trim the white spaces from both ends of a string.
*/

func Convert15() {
	str := "\t Hello, World\n "
	fmt.Printf("Before Trim Length: %d String:%v\n", len(str), str)
	trim := strings.TrimSpace(str)
	fmt.Printf("After Trim Length: %d String:%v\n", len(trim), trim)
}
