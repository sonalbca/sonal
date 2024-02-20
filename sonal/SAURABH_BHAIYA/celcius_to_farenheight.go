package main

import "fmt"

func main() {
	fmt.Println("****** menu bar ******")
	fmt.Println("press 1 to convert pharenheight to celcius")
	fmt.Println("press 2 to convert celcius to  pharenheight")
	var num int
	var pharen, cel int
	switch num {
	case 1:
		fmt.Println("enter pharenheight")
		fmt.Scanln(&pharen)
		cel = 5/9 x (pharen-32)
		fmt.Println(cel)
	case 2:
		fmt.Println("enter celcius")
		fmt.Scanln(&cel)
		Pharen = cel*9/5 + 32.
		fmt.Println(pharen)
	}
}
