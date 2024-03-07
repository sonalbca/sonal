package main

import (
	"encoding/json"
	"fmt"
)

type MaritalDt struct {
	Name          string
	MaritalStatus string
	NameFm        string
	Kids          int
}
type SaurabhFm struct {
	FM        int
	FmDetails []string
	MrDt      []MaritalDt
}

func main() {
	jsonData := &SaurabhFm{
		FM:        5,
		FmDetails: []string{"Gautam kumar singh", "Binita Devi", "Saurabh kumar", "Rahul kumar", "Raja"},

		MrDt: []MaritalDt{
			{
				Name:          "saurabh",
				MaritalStatus: "unmarried",
				NameFm:        "No family",
				Kids:          0,
			},
			{
				Name:          "Rahul",
				MaritalStatus: "unmarried",
				NameFm:        "No family",
				Kids:          0,
			},
			{
				Name:          "Raja",
				MaritalStatus: "unmarried",
				NameFm:        "No family",
				Kids:          0,
			},
		},
	}

	j, er := json.Marshal(jsonData)
	if er != nil {
		fmt.Println("could not unmarshal")
	}
	fmt.Println("json ", string(j))
}
