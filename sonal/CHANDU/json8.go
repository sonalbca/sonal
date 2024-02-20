package main

import (
	"encoding/json"
	"fmt"
)

type Familymember struct {
	FathersName     string   `json:"fathers_name"`
	MothersName     string   `json:"mothers_name"`
	BrothersName    string   `json:"brothers_name"`
	BrothersDetails []string `json:"brothers_details"`
}

func main() {
	fm := Familymember{}
	fmt.Println("enter your fathers name")
	fmt.Scanln(&fm.FathersName)
	fmt.Println("enter your mothers name")
	fmt.Scanln(&fm.MothersName)
	fmt.Println("enter your brothers name")
	fmt.Scanln(&fm.BrothersName)
	BD := []string{"name", "naman kumar", "class", "bed"}
	jsonstring, _ := json.Marshal(BD)
	fmt.Println(string(jsonstring))

}
