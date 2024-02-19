package main

import "fmt"

func TargetSum(array []int, target int) []int {
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
	array := []int{2, 7, 11, 15}
	sl := TargetSum(array, 17)
	fmt.Println(sl)
}
