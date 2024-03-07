package main

import "fmt"

type PEerson struct {
	name     string
	village  string
	mobileno int
	pincode  int
}
type customerDetails struct {
	Acount_number int
	IFC_code      int
	Amount        int
	PE            PEerson
}

func main() {
	var DoesItExists = map[string]bool{}
	DoesItExists["DMS1001"] = true // MYSQL SQL DATABASE
	DoesItExists["DMS1002"] = true
	DoesItExists["DMS1003"] = true
	DoesItExists["DMS1004"] = true

	Bank := map[string]customerDetails{}
	Bank["DMS1001"] = customerDetails{
		Acount_number: 12345, IFC_code: 100, Amount: 1000, PE: PEerson{name: "SAURABH", village: "AURANGABAD", mobileno: 9334379889, pincode: 824102},
	}
	Bank["DMS1002"] = customerDetails{
		Acount_number: 23456, IFC_code: 101, Amount: 1001, PE: PEerson{name: "PARAS", village: "DEV", mobileno: 9693772664, pincode: 824101},
	}
	Bank["DMS1003"] = customerDetails{
		Acount_number: 23456, IFC_code: 101, Amount: 1001, PE: PEerson{name: "HARSHIT", village: "MOKHTARPUR", mobileno: 9693772664, pincode: 824103},
	}
	Bank["DMS1004"] = customerDetails{
		Acount_number: 23456, IFC_code: 101, Amount: 1001, PE: PEerson{name: "SONAL", village: "BUMARU", mobileno: 9608023850, pincode: 824102},
	}
	var userid string
	var mobileno string
	fmt.Println("enter your userid")
	fmt.Scanln(&userid)
	fmt.Println("enter your mobile no")
	fmt.Scanln(&mobileno)
	if DoesItExists[userid] {
		fmt.Println("Customer Exists and their Data is Below")
		fmt.Println(Bank[userid])
	} else {
		fmt.Println("Customer Does not Exists")
	}

}
