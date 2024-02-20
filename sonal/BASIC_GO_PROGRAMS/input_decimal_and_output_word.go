package main

import "fmt"

func findlength(num int) int {
	count := 0
	for num > 0 {
		count++
		num = num / 10
	}
	return count
}
func main() {
	var num int
	fmt.Println("enter the number")
	fmt.Scanln(&num)
	n := findlength(num)
	fmt.Println(num)
	fmt.Print(n)
	switch n {
	case 1:
		callforone(num)
	case 2:
		callforTwo(num)
	case 3:
		callforThree(num)

	}

}
func callforThree(num int) {
	rem := num % 100
	num = num / 100
	callforone(num)
	fmt.Print(" Hundred")
	callforTwo(rem)
}
func callforTwo(num int) {
	if num/10 == 1 {
		switch num {
		case 10:
			fmt.Print(" Ten")
		case 11:
			fmt.Print(" Eleven")
		case 12:
			fmt.Print(" Twelve")
		case 13:
			fmt.Print(" Thirteen")
		case 14:
			fmt.Print(" Forteen")
		case 15:
			fmt.Print(" Fifteen")
		case 16:
			fmt.Print(" Sixteen")
		case 17:
			fmt.Print(" Seventeen")
		case 18:
			fmt.Print(" Eighteen")
		case 19:
			fmt.Print(" Nineteen")
		default:

		}
	} else {
		rem := num % 10
		num = num / 10
		switch num {
		case 2:
			fmt.Print(" Twenty")
			callforone(rem)
		case 3:
			fmt.Print(" Thirty")
			callforone(rem)
		case 4:
			fmt.Print(" Forty")
			callforone(rem)
		case 5:
			fmt.Print(" Fifty")
			callforone(rem)
		case 6:
			fmt.Print(" Sixty")
			callforone(rem)
		case 7:
			fmt.Print(" Seventy")
			callforone(rem)
		case 8:
			fmt.Print(" Eighty")
			callforone(rem)
		case 9:
			fmt.Print(" Ninety")
			callforone(rem)
		default:
			callforone(rem)
		}
	}
}

func callforone(num int) {
	switch num {
	case 1:
		fmt.Print(" One")
	case 2:
		fmt.Print(" Two")
	case 3:
		fmt.Print(" Three")
	case 4:
		fmt.Print(" Four")
	case 5:
		fmt.Print(" Five")
	case 6:
		fmt.Print(" Six")
	case 7:
		fmt.Print(" Seven")
	case 8:
		fmt.Print(" Eight")
	case 9:
		fmt.Print(" Nine")
	default:
		fmt.Print(" Zero")
	}
}
