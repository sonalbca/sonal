package main

import "fmt"

type st struct {
	vill string
	mob  int
}
type Person struct {
	BankAcc     int
	name        string
	DOB         string
	father_name string
	address     st
}

func main() {
	var employee = map[string]int{"Saurabh": 1, "Sonal": 2, "Paras": 3}
	fmt.Println(employee["Paras"])
	employee["Harshit"] = 4
	employee["Asmriti"] = 5

	var DoesItExists = map[string]bool{}
	DoesItExists["DMS012221"] = true // MYSQL SQL DATABASE
	DoesItExists["DMS013331"] = true

	var Emp = map[string]Person{}
	Emp["DMS012221"] = Person{BankAcc: 235900150003420, name: "Saurabh", DOB: "20/03/2001", father_name: "Gautam Kr Singh", address: st{vill: "Belsara", mob: 9334379889}}
	Emp["DMS013331"] = Person{BankAcc: 139879874917384, name: "Harsit", DOB: "25/02/2006", father_name: "Jay Shankar Prasad", address: st{vill: "Mokhityar Pur", mob: 83784389}}

	var userId string
	fmt.Println("Serach User ")
	fmt.Scanln(&userId)
	if DoesItExists[userId] {
		fmt.Println("Customer Exists and their Data is Below") // 202 --- http request successfull
		fmt.Println(Emp[userId])
	} else {

		fmt.Println("Customer Does Not Exist")
	}
}
