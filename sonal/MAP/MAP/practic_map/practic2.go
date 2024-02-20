package main

import "fmt"

func sortInteger(s []int) {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s)-1; j++ {
			if s[i] > s[j] {
				temp := s[i]
				s[i] = s[j]
				s[j] = temp
			}
		}
	}

}

func main() {
	var map1 = map[int]string{}
	map1[1] = "kailas"
	map1[3] = "nath"
	map1[2] = "pati"
	map1[4] = "ki"
	map1[5] = "jay"
	fmt.Println(map1)

	slice := []int{}
	for k, _ := range map1 {
		slice = append(slice, k)
	}
	fmt.Println(slice)
	sortInteger(slice)
	fmt.Println(slice)
	for key, _ := range slice {
		fmt.Println(key, " ", map1[key])

	}
}
