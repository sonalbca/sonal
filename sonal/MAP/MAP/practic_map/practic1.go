package main

import "fmt"

func sortinteger(s []int) {
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
	map1[1] = "Ram"
	map1[12] = "shayam"
	map1[3] = "kajal"
	map1[23] = "keshav"
	fmt.Println(map1)
	fmt.Println(map1[2])

	var map2 = map[string]int{"sonal": 1, "shayam": 2, "rahul": 3}
	fmt.Println(map2)
	fmt.Println(map2["shayam"])

	map3 := make(map[string]int)
	map3["roll"] = 1
	map3["class"] = 4
	fmt.Println(map3)

	slice := []int{}
	for k, _ := range map1 {
		slice = append(slice, k)
	}
	fmt.Println(slice)
	sortinteger(slice)
	fmt.Println(slice)
	for _, key := range slice {
		fmt.Println(key, " ", map1[key])
	}
}
