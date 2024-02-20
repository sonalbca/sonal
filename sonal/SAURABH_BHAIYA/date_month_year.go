package main

import (
	"fmt"
	"strconv"
)

func findFinalDay(f int) {
	switch f {
	case 0:
		fmt.Println("Sunday")
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 7:
		fmt.Println("Saturday")
	}
}

func LeapYearOrNot(year int) bool {
	if year%4 == 0 {
		return true
	}
	return false
}

// Get the year code
func findYearCode(year int) int {
	// get the last two digit of your year 2023  28 /7 4 ==0
	year = year % 100
	code := (year + (year / 4)) % 7
	return code
}

// Get century code
/*
1700s = 4
1800s = 2
1900s = 0
2000s = 6
2100s = 4
2200s = 2
2300s = 0

2024 ==
rem = 2024%100 == 24
2024-rem = 2024- 24 == 2000
*/
func centuryCode(year int) int {
	rem := year % 100
	cCode := year - rem
	centuryMap := map[int]int{}
	centuryMap[1700] = 4
	centuryMap[1800] = 2
	centuryMap[1900] = 0
	centuryMap[2000] = 6
	centuryMap[2100] = 4
	centuryMap[2200] = 2
	centuryMap[2300] = 0

	centureCodeF := centuryMap[cCode]
	return centureCodeF
}
func main() {
	var date string
	fmt.Println("Enter the date")
	fmt.Scanln(&date)
	day := ""
	month := ""
	year := ""
	count := 0
	for _, i := range date {
		if string(i) == "/" {
			count++
		} else {
			if count == 0 {
				day = day + string(i)
			} else if count == 1 {
				month = month + string(i)
			} else {
				year = year + string(i)
			}
		}
	}

	fmt.Println("day =", day)
	fmt.Println("Month =", month)
	fmt.Println("Year =", year)
	// month code
	//(Year Code + Month Code + Century Code + Date Number - Leap Year Code) mod 7
	// finding year code
	intYear, _ := strconv.Atoi(year)
	yearCode := findYearCode(intYear)
	fmt.Println(yearCode)
	//January = 0
	//February = 3
	//March = 3
	//April = 6
	//May = 1
	//June = 4
	//July = 6
	//August = 2
	//September = 5
	//October = 0
	//November = 3
	//December = 5

	monthCode := []int{0, 3, 3, 6, 1, 4, 6, 2, 5, 0, 3, 5}
	CentCode := centuryCode(intYear)
	//fmt.Println(CentCode)
	isLeapYear := LeapYearOrNot(intYear)
	//(Year Code + Month Code + Century Code + Date Number - Leap Year Code) mod 7
	intMon, _ := strconv.Atoi(month)
	intDate, _ := strconv.Atoi(day)
	if isLeapYear {
		//2 +
		f := (yearCode + monthCode[intMon-1] + CentCode + intDate - 1) % 7
		findFinalDay(f)
	} else {
		f := (yearCode + monthCode[intMon-1] + CentCode + intDate - 0) % 7

		findFinalDay(f)
	}
}
