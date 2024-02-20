package main

import "fmt"

func linearsearch(Array []int, key int) bool {
	for _, item := range Array {
		if item == key {
			return true
		}
	}
	return false
}

func main() {
	Array := []int{95, 78, 46, 58, 45, 86, 99, 251, 320}
	fmt.Println(linearsearch(Array, 58))
}
