package main

import "fmt"

func TargetArr(Array []int, sum int) []int {

	Sl := []int{}

	for i := 0; i < len(Array); i++ {
		for j := i + 1; j < len(Array); j++ {

			if Array[i]+Array[j] == sum {

				Sl = append(Sl, i)
				Sl = append(Sl, j)
				return Sl

			}
		}
	}
	return nil
}
func main() {
	Array := []int{8, 7, 6, 5, 4, 2}
	Sl := TargetArr(Array, 9)
	fmt.Println(Sl)
}
