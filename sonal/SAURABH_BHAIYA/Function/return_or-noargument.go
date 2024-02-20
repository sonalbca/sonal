package main

import "fmt"

func find() (int, float64, string, bool) {
	a := 67.9
	return int(a), a, "saurabh", true
}
func main() {
	fmt.Println(find())
}
