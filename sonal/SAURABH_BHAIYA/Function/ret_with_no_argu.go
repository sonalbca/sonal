package main

import "fmt"

func per(Gross int, pt int) int {
	pert := (Gross * pt) / 100
	return pert
}

// function with argument and with return
func TotalSal(anil int, hra int, pf int, med int) int {
	Total := anil + per(anil, hra) + per(anil, pf) + per(anil, med)
	return Total
}

func main() {

	Anil := 40000
	hra := 2
	pf := 7
	med := 5
	total := TotalSal(Anil, hra, pf, med)
	fmt.Println("Anil total Salary = ", total)

}
