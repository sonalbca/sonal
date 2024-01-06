package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var s = "100"
	intvar, err := strconv.Atoi(s)

	//here the type of
	fmt.Println(intvar, err, reflect.TypeOf(intvar))
	
	//here convert string to integer
	sk, _ := strconv.Atoi(s)
	fmt.Println(sk)
}
