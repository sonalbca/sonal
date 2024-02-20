package main

import "fmt"

func main() {
	var NestedMap = map[string]map[string]string{}
	NestedMap["name"] = map[string]string{}
	//NestedMap["BankAC"] = map[string]string{}

	NestedMap["Name"]["a"] = "sonal"
	NestedMap["name"]["b"] = "BankAC"

	NestedMap["BankAC"]["bankAc"] = "12345678910"
	NestedMap["bankAC"]["Amount"] = "654012"

	for key, val := range NestedMap {
		fmt.Println(key, val)
	}
}
