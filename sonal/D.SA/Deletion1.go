package main

import "fmt"

func IsAvailable(array []int, del int) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == del {
			return true
		}
	}
	return false
}
func Deletion(array []int, del int) []int {
	sl := make([]int, 0, len(array)-1)
	isAvl := IsAvailable(array, del)
	for isAvl {
		for i := 0; i < len(array); {
			if array[i] != del {
				sl = append(sl, array[i])
				i++

			} else {
				i++
			}

		}
		return sl
	}
	return nil
}
func main() {
	array := []int{12, 15, 18, 5}
	sl := Deletion(array, 12)
	fmt.Println("original array", array)
	if sl != nil {
		fmt.Println("after deleting element", sl)
	} else {
		fmt.Println("element not found")
	}
}
