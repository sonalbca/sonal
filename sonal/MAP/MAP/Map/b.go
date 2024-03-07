package main

import "fmt"

type details struct {
	name      string
	acount_no int
	balance   int
}

func Dumidata(userid string) (details, bool) {
	customerDetails := map[string]details{}
	customerDetails["dms100"] = details{name: "sonal", acount_no: 1000000, balance: 43200}
	customerDetails["dms101"] = details{name: "paras", acount_no: 1000001, balance: 43234}
	customerDetails["dms102"] = details{name: "harshit", acount_no: 1000002, balance: 4325}
	customerDetails["dms103"] = details{name: "patel", acount_no: 1000003, balance: 43206}
	data, ok := customerDetails[userid]
	if ok {
		return data, true
	} else {
		fmt.Println("No Data")
		return data, false
	}
}

func main() {
	pnb := map[string]details{}
	fmt.Println("enter user id")
	var userid string
	fmt.Scanln(&userid)
	data, ok := Dumidata(userid)
	if ok {
		fmt.Println("your balance is ", data.balance)
		fmt.Println("enter withdrwo amount")
		var withdro int
		fmt.Scanln(&withdro)
		if withdro <= data.balance {
			data.balance = data.balance - withdro
			pnb[userid] = data
			fmt.Println("Remaining Balance:", data.balance)
		} else {
			fmt.Println("You do not have Sufficient balance.....enter laser amount")
		}
	} else {
		fmt.Println("No Data User Available")
	}
}
