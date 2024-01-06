package main

import "fmt"

func findMedianSortedArrays(arr []int, arr1 []int) float64 {
	firstArrayLen := len(arr)
	secondArrayLen := len(arr1)

	var mid int
	i := 0
	j := 0
	var k int
	mid = (firstArrayLen+secondArrayLen)/2 + 1

	//This is the case in which the total lenght of two arrays is odd and there is only one median
	if (firstArrayLen+secondArrayLen)%2 == 1 {
		var median float64

		for k < mid {
			if i < firstArrayLen && j < secondArrayLen {
				if arr[i] <= arr1[j] {
					median = float64(arr[i])
					i++
					k++
				} else {
					median = float64(arr1[j])
					j++
					k++
				}
			} else if i < firstArrayLen {
				median = float64(arr[i])
				i++
				k++
			} else {
				median = float64(arr1[j])
				j++
				k++
			}

		}
		return median
	} else { //This is the case in which the total lenght of two arrays is even and there is only two medians. We need to return average of these two medians
		var median1 float64
		var median2 float64

		for k < mid {
			median1 = median2
			if i < firstArrayLen && j < secondArrayLen {
				if arr[i] <= arr1[j] {
					median2 = float64(arr[i])
					i++
					k++
				} else {
					median2 = float64(arr1[j])
					j++
					k++
				}
			} else if i < firstArrayLen {
				median2 = float64(arr[i])
				i++
				k++
			} else {
				median2 = float64(arr1[j])
				j++
				k++
			}

		}
		return (median1 + median2) / 2
	}
}
func main() {
	arr := []int{1, 3}
	arr1 := []int{2}
	median := findMedianSortedArrays(arr, arr1)
	fmt.Println(median)
}
