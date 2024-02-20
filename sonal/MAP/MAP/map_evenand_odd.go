package main

import "fmt"

func EvenAndOdd() {
	var EvenOdd = make(map[string]int)
	EvenOdd["sonal"] = 1
	EvenOdd["harshit"] = 2
	EvenOdd["saurabh"] = 3
	EvenOdd["asmirity"] = 4
	EvenOdd["kajal"] = 5
	EvenOdd["chandu"] = 6
	EvenOdd["kala"] = 7
	EvenOdd["gora"] = 8
	EvenOdd["ab hm"] = 9
	EvenOdd["kya bole"] = 10
	for key, val := range EvenOdd {
		if val%2 == 0 {
			fmt.Println(key, val)
		}
	}
}
