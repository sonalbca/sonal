package main

import "fmt"

type PErson struct {
	name  string
	class string
	roll  int
	vill  string
}

func main() {
	PDetails := map[string]PErson{}
	PDetails["DETAILS"] = PErson{name: "sonal", class: "bca 2nd year", vill: "BUMARU", roll: 83}
	fmt.Println(PDetails)
	_, ok := PDetails["DETAILS"]
	if ok {
		count := 0
		for k, _ := range PDetails {
			if k == "DETAILS" {
				PDetails["DETAILS"] = PErson{name: "SAURABH", class: "bca final year", vill: "Dev belsara", roll: 1}
			}
			count++
		}
		fmt.Println(count)
	}
	fmt.Println(PDetails)
}
