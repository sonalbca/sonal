package main

import "fmt"

func main() {
	Slice1 := make([]int, 2)
	Slice1[0] = 15
	Slice1[1] = 25
	Slice1 = append(Slice1, 20, 30, 40, 50)
	fmt.Println(Slice1)
	//change item
	Slice1[0] = 5
	Slice1[2] = 12
	fmt.Println(Slice1)
}
