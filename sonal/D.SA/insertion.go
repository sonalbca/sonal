package main

import "fmt"

func main() {
	var array = make([]int, 0, 6)
	array = append(array, 2, 5, 3, 7, 6)
	for i := 0; i < len(array); i++ {
		fmt.Print(" ", array[i])
	}
	fmt.Println()
	var index = 3
	array[index] = 10
	for i := 0; i < len(array); i++ {
		fmt.Print(" ", array[i])
	}

}
