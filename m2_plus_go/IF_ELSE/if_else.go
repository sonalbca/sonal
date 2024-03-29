package main

import "fmt"

func main() {
	Ifelse1()
	//If_else2()
	//If_else3()
	//If_else4()
}

//----------------- 1 ------------------------//
/*
If Else
In this tutorial you'll learn how to write decision-making conditional statements used to perform different actions in Golang.

Golang Conditional Statements
Like most programming languages, Golang borrows several of its control flow syntax from the C-family of languages. In Golang we have the following conditional statements:

The if statement - executes some code if one condition is true
The if...else statement - executes some code if a condition is true and another code if that condition is false
The if...else if....else statement - executes different codes for more than two conditions
The switch...case statement - selects one of many blocks of code to be executed
We will explore each of these statements in the coming sections.

Golang - if Statement
The if statement is used to execute a block of code only if the specified condition evaluates to true.

Syntax
if  condition {
// code to be executed if condition is true
}
The example below will output "Japan" if the X is true:
*/

func Ifelse1() {
	var s = "Japan"
	x := true
	if x {
		fmt.Println(s)
	}
}

//---------------------- 2 ----------------------
/*
Golang - if...else Statement
The if....else statement allows you to execute one block of code if the specified condition is evaluates to true and another block of code if it is evaluates to false.

Syntax
if  condition {
// code to be executed if condition is true
} else {
// code to be executed if condition is false
}
The example below will output "Japan" if the X is 100:
*/

func If_else2() {
	x := 100

	if x == 100 {
		fmt.Println("Japan")
	} else {
		fmt.Println("Canada")
	}
}

//--------------------- 3 ------------------------
/*
Golang - if...else if...else Statement
The if...else if...else statement allows to combine multiple if...else statements.

Syntax
if  condition-1 {
// code to be executed if condition-1 is true
} else if condition-2 {
// code to be executed if condition-2 is true
} else {
// code to be executed if both condition1 and condition2 are false
}
The example below will output "Japan" if the X is 100:
*/

func If_else3() {
	x := 100

	if x == 50 {
		fmt.Println("Germany")
	} else if x == 100 {
		fmt.Println("Japan")
	} else {
		fmt.Println("Canada")
	}
}

//-------------------- 4 ------------------
/*
Golang - if statement initialization
The if statement supports a composite syntax where the tested expression is preceded by an initialization statement.

Syntax
if  var declaration;  condition {
// code to be executed if condition is true
}
The example below will output "Germany" if the X is 100:
*/

func If_else4() {
	if x := 100; x == 100 {
		fmt.Println("Germany")
	}
}
