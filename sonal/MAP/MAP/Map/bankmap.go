package main

import (
	"fmt"
)

type Details struct {
	AccNo   int
	name    string
	dob     string
	balance int
}

func dummyData(userId string) (Details, bool) {
	CDetails := map[string]Details{}
	CDetails["@saurabh"] = Details{AccNo: 112333, name: "saurabh", dob: "20/03/2001", balance: 83834}
	CDetails["@harshit"] = Details{AccNo: 1397974, name: "harshit", dob: "20/03/2004", balance: 83833}
	CDetails["@sonal"] = Details{AccNo: 2973434, name: "sonal", dob: "20/03/2004", balance: 3443444}
	CDetails["@asmriti"] = Details{AccNo: 1397974, name: "Asmriti", dob: "20/03/2005", balance: 386123}

	data, ok := CDetails[userId]
	if ok {
		return data, true
	} else {
		fmt.Println("No Data")
		return data, false
	}
}

func main() {
	pnb := make(map[string]Details)
	fmt.Println("enter the UserId")
	var userId string
	fmt.Scan(&userId)
	data, ok := dummyData(userId)

	if ok {
		fmt.Println("Your Balance is: ", data.balance)
		fmt.Println("Enter withdrow data")
		var withdrow int
		fmt.Scan(&withdrow)
		if withdrow <= data.balance {
			fmt.Println("Collect your money")
			data.balance = data.balance - withdrow
			pnb[userId] = data
			fmt.Println("Remaining Balance:", data.balance)
		} else {
			fmt.Println("You do not have Sufficient balance.....eneter laser amount")
		}
	} else {
		fmt.Println("No Data User Available")
	}
}
