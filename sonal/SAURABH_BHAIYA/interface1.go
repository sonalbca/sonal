package main

import "fmt"

type Emplyoee interface {
	PrinttName() // method
	PrintRoll()
}

func (e emp) PrinttName() {
	fmt.Println("Sonal babu kaisan ba ")
}

func (e emp) PrintRoll() {
	fmt.Println("Harshit babu tum 1000 ")
}

type emp string // emp new type

func main() {
	var eployes Emplyoee // interface intatitiation

	eployes = emp("") // type use

	eployes.PrinttName()
	eployes.PrintRoll()

}
