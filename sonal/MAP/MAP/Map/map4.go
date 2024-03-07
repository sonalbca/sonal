package main

import "fmt"

func main() {
	Class := map[string]int{}
	Class["BCA"] = 83
	Class["B.sc"] = 12
	Class["Mca"] = 19
	Class["B.com"] = 14
	fmt.Println("Before class")
	fmt.Println(Class)

	Class["B.com"] = 102
	fmt.Println("After class")
	fmt.Println(Class)
	fmt.Println("Delete class")
	delete(Class, "Mca")
	fmt.Println(Class)
}
