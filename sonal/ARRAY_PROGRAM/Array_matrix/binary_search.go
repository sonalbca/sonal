package main

import "fmt"

//binary search
// 12 23 24 25 67 69 sorted binary search is always perform on sorted Array

func BinSear(arr []int, ser int) int {
	si := 0
	ei := len(arr) - 1
	mid := 0

	for si < ei {
		mid = (si + ei) / 2

		if arr[mid] == ser {
			return mid
		} else if arr[mid] < ser {
			si = mid + 1
		} else {
			ei = mid - 1
		}

	}
	return 0
}

func main() {
	arr := []int{12, 34, 65, 87, 99}
	var ser = 34
	s := BinSear(arr, ser)
	fmt.Println("original array", arr)
	if s != ser {
		fmt.Println("element present at index", s, ser)
	} else {
		fmt.Println(ser, "element not found")
	}
}
