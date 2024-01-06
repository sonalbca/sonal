package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	//BooleanToString()
	BToS()
}
func BooleanToString() {
	var b bool = true
	fmt.Println(reflect.TypeOf(b))

	var s string = strconv.FormatBool(true)
	fmt.Println(reflect.TypeOf(s))
}
func BToS() {
	b := true
	s := fmt.Sprintf("%v", b)
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))
}
