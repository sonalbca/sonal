package main

import "fmt"

func main() {
	m := make(map[string]int) // make se aap empty map create kr sakte hai
	// elemnt adding in empty map
	m["a"] = 3
	m["b"] = 4
	m["c"] = 5
	m["d"] = 6
	m["e"] = 7
	m["f"] = 8
	m["g"] = 9
	m["h"] = 10

	// userId ==a==80
	fmt.Println(m)

	_, ok := m["e"]
	if ok {
		count := 0
		for k, _ := range m {
			if k == "e" {
				m[k] = 800
			}
			count++
		}
		fmt.Println(count)
	}
	fmt.Println(m)
}
