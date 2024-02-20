package main

import "fmt"

func main() {
	Array1 := [6]int{2, 13, 14, 5, 67, 89}
	for i := 0; i < 6; i++ {
		fmt.Println(Array1[i])
	}
	i := 2
	if i < 0 || i >= len(Array1) {
		fmt.Println("the given index is out of bound")

	} else {
		newlength := 0
		for index := range Array1 {
			if i != index {
				Array1[newlength] = Array1[index]
				newlength++
			}
		}
		newArray := Array1[:newlength]
		fmt.Println("the new array is ", newArray)
	}
}
