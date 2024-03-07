package main

import "fmt"

type Sonal struct {
	name          string
	kids          int
	MaritalStatus bool
}
type FamilyMember struct {
	fm           int
	Sonaldetatis []Sonal
}

func main() {
	F := FamilyMember{}
	F.fm = 14
	F.Sonaldetatis = []Sonal{
		{
			name:          "sonal kumar singh",
			kids:          2,
			MaritalStatus: true,
		},
		{
			name:          "kundan kumar singh",
			kids:          4,
			MaritalStatus: true,
		},
		{
			name:          "naman kumar singh",
			kids:          3,
			MaritalStatus: true,
		},
	}
	fmt.Println("family member", F.fm)
	for _, val := range F.Sonaldetatis {
		fmt.Println(val)
	}

}
