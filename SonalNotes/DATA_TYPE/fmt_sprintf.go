package main

import (
	"fmt"
	"reflect"
)

// string from integer
func main() {
	b := 1225
	fmt.Println(reflect.TypeOf(b))

	s := fmt.Sprintf("%v", b)
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))
}
