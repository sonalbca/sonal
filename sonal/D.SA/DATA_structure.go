package main

import (
	"fmt"
)

func arraysimple(array [5]int, size int) int {
	avg := 0
	sum := 0
	count := 0
	for i := 0; i < size; i++ {
		sum = sum + array[i]
		count++
		avg = sum / count
	}
	fmt.Println("average of all digit is ", avg)
	return sum
}
func main() {
	array := [5]int{12, 13, 14, 15, 16}
	var res int

	res = arraysimple(array, 5)
	fmt.Println("sum of all digit is ", res)
}
