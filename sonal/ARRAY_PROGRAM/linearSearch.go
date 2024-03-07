package main

import "fmt"

func lishr(Array []int, find int) bool {
	for _, item := range Array {
		if item == find {
			return true
		}
	}
	return false
}

func main() {
	Array := []int{95, 78, 46, 58, 45, 86, 99, 251, 320}
	fmt.Println(lishr(Array, 58))
}
