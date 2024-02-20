package main

import "fmt"

func RecTangle(l int, b int) (area int, parimeter int) {
	parimeter = 2 * (l * b)
	area = l * b
	return
}
func main() {
	var l, b int
	l, b = RecTangle(6, 5)
	fmt.Println("Area", l)
	fmt.Println("Parimeter", b)
}
