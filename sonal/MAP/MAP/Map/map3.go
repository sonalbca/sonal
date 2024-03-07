package main

import "fmt"

func Isemp(EMP map[string]int) {
	if len(EMP) == 0 {
		fmt.Println("map is empty")
	} else {
		fmt.Println("map len is ", len(EMP))
	}
}

func main() {
	var EMP = map[string]int{}
	EMP["saurabh"] = 12
	EMP["sonal"] = 13
	EMP["paras"] = 14
	EMP["harshit"] = 15
	EMP["asmirity"] = 16
	EMP["kajal"] = 17
	Isemp(EMP)

}
