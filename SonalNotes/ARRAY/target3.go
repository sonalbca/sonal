package main

import "fmt"

func Target3(array []int, target int) []int {
	sl := []int{}
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {

			if array[i]+array[j] == target {
				sl = append(sl, i)
				sl = append(sl, j)
				return sl
			}
		}
	}
	return nil
}
func main() {
	array := []int{3, 3}
	sl := Target3(array, 6)
	fmt.Println(sl)
}
