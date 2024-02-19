package main

import (
	"errors"
	"fmt"
	"log"
)

func BinarySearch(arr []int, ser int) (int error) {
	si := 0
	ei := len(arr) - 1

	for si < ei {
		mid := (si + ei) / 2

		if arr[mid] == ser {
			return mid, nil

		} else if arr[mid] < ser {
			si = mid + 1
		} else {
			ei = mid - 1
		}
	}
	return 0, errors.New("element not found")
}
func main() {
	arr := []int{2, 3, 4, 5, 6, 7, 8}
	value, err := BinarySearch(arr, 100)
	if err == nil {
		fmt.Println("element present at index", value)
	} else {
		log.Fatal(err)
	}

}
