package main

import "fmt"

func SimpleInterest(p float32, r float32, t float32) float32 {
	si := p * r * t / 100
	fmt.Print(si)
	return si
}
func main() {
	var p, r, t float32
	fmt.Println("enter principle")
	fmt.Scanln(&p)
	fmt.Println("enter rate")
	fmt.Scanln(&r)
	fmt.Println("enter time")
	fmt.Scanln(&t)
	SimpleInterest(p, r, t)
}
