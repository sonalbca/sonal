package main

//check whether an alphabet is vowel or consonant using switch case
import "fmt"

func main() {
	var character string
	fmt.Println("enter your name")
	fmt.Scanln(&character)
	switch character {
	case "a":
		fmt.Println("vowel")
	case "e":
		fmt.Println("vowel")
	case "i":
		fmt.Println("vowel")
	case "o":
		fmt.Println("vowel")
	case "u":
		fmt.Println("vowel")
	default:
		fmt.Println("consonant")
	}

}
