package main

import "fmt"

type SStudent struct {
	name  string
	roll  int
	phone int
}

func main() {
	SS := SStudent{}
	SS.name = "Sam"
	SS.roll = 12
	SS.phone = 12345678910
	fmt.Println(SS)

	SS1 := SStudent{}
	SS1.name = "john"
	SS1.roll = 13
	SS1.phone = 10987654321
	fmt.Println(SS1)
}
