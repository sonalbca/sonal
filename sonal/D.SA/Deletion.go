package main

import "fmt"

func main() {
	array := []int{2, 7, 4, 5, 9}
	fmt.Println("original array", array)
	del := 7
	sl := make([]int, len(array)-1)

	k := 0
	for i := 0; i < len(array); {
		if array[i] == del {
			i++
		} else {
			sl[k] = array[i]
			k++
			i++
		}
	}
	fmt.Println("after deleting element", sl)
}
