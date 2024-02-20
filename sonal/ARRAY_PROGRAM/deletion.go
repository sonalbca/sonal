package main

import "fmt"

func isElementAvailable(array [5]int, del int) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == del {
			return true
		}
	}
	return false
}
func deleteElement(array [5]int, del int) []int { // sl = 12 , 23,
	// is element availabe in given array
	sl := make([]int, 0, len(array)-1)        // making a slice
	isAvail := isElementAvailable(array, del) // checking for element available in array or not
	if isAvail {
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
	array := [5]int{12, 23, 24, 54, 66}
	sl := deleteElement(array, 23)

	fmt.Println("original array =", array)
	if sl != nil {
		fmt.Println("after deleting array is ", sl)
	} else {
		fmt.Println("elemenet not found")
	}
}
