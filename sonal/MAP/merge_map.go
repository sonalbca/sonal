package main

import "fmt"

func main() {
	map1 := map[string]int{"raju": 12, "kaju": 2, "mahesi": 12, "basanti": 115}
	map2 := map[string]int{"ram": 17, "monday": 41, "kalh": 18}

	for key, val := range map2 {
		map1[key] = val
	}
	fmt.Println(map1)
}
