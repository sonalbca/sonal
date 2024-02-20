package main

import "fmt"

func main() {
	a := [5]int{12, 23, 44, 54, 5}
	fmt.Println("o a", a)

	del := 23
	arr := make([]int, len(a)-1)
	k := 0
	for i := 0; i < len(a); {
		if a[i] == del {
			i++
		} else {
			arr[k] = a[i]
			i++
			k++
		}
	}
	fmt.Println("d a", arr)
}
