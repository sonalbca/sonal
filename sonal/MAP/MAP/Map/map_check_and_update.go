package main

import "fmt"

func main() {

	var DoesItExists = map[string]bool{}
	DoesItExists["SAURABH"] = true // MYSQL SQL DATABASE
	DoesItExists["SONAL"] = true
	DoesItExists["PARAS"] = true
	DoesItExists["HARSHIT"] = true
	DoesItExists["ASMIRITY"] = true

	MAP := map[string]int{}
	MAP["SAURABH"] = 12
	MAP["SONAL"] = 13
	MAP["PARAS"] = 14
	MAP["HARSHIT"] = 15
	MAP["ASMIRITY"] = 16
	var userId string
	fmt.Println("Serach User ")
	fmt.Scanln(&userId)
	if DoesItExists[userId] {
		fmt.Println(" Data is Below")
		fmt.Println(MAP[userId])
	} else {

		fmt.Println("Data Does Not Exist")

	}
	MAP["SONAL"] = 225
	fmt.Println(MAP)
}
