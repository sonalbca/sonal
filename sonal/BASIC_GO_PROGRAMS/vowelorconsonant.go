package main

import "fmt"

func main() {
	//VowelAndConsonant()
	Vowelconsonant()
}

func VowelAndConsonant() {
	var name string
	fmt.Println("enter your name")
	fmt.Scanln(&name)
	if name == "a" || name == "e" || name == "i" || name == "o" || name == "u" {
		fmt.Println("vowel", name)
	} else if name == "A" || name == "E" || name == "I" || name == "O" || name == "U" {
		fmt.Println("vowel", name)
	} else {
		fmt.Println("consonant", name)
	}
}
func Vowelconsonant() {
	var name string
	fmt.Println("enter your name")
	fmt.Scanln(&name)
	for i := 0; i < len(name); i++ {
		if name == "a" || name == "e" || name == "i" || name == "o" || name == "u" {
			fmt.Println("vowel", name)
		} else if name == "A" || name == "E" || name == "I" || name == "O" || name == "u" {
			fmt.Println("vowel", name)
		} else {
			fmt.Println("consonant", name)
		}
	}
}
