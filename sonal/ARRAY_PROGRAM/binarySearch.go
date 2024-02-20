package main

import (
	"errors"
	"fmt"
	"log"
)

//binary search
// 12 23 24 25 67 69 sorted binary search is always perform on sorted Array

func BinSear(arr []int, ser int) (int, error) {
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
	arr := []int{12, 34, 65, 87, 99}

	value, err := BinSear(arr, 100)
	if err == nil {
		fmt.Println("element present at index", value)
	} else {
		log.Fatal(err)
	}

}
