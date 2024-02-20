package main

import "fmt"

func main() {
	//loop1() //ans not good
	//loop2()
	//loop3()
	//loop4()
	//loop5()
	//loop6()
	//loop7()
	//loop8()
	//loop9()
	//loop10()
	//loop11() //not proper work
	//loop12()
	//loop13()
	//EvenAndOdd()
	//EvenAndOdd1()
	//PrintUserInput()
	//pattern()
	//loop14()
	//loop15()
	//loop16()
	//loop17()
	//loop18()

}

func loop1() {
	var n int
	fmt.Println("enter the row")
	fmt.Scanln(&n)
	for i := 0; i <= n; i++ { //row
		for j := 0; j < 2*n; j++ {
			if j%2 == 0 {
				fmt.Print("*")
			} else {
				fmt.Print("**")
			}

		}
		fmt.Println()
	}
	/*

	 */

}
func loop2() {
	var n int
	fmt.Println("enter the row")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		for sp := 1; sp <= n*2-i*2; sp++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 1; i <= n-1; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Print("*")
		}
		for sp := 1; sp <= i*2; sp++ {
			fmt.Print(" ")
		}
		for j := 1; j <= n-i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	/*
						*      *
						**    **
						***  ***
						********
						***  ***
		    			**    **
		    			*      *

	*/
}
func loop3() {
	for i := 0; i < 4; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(" ", i+1)
		}
		fmt.Println()
	}

	/*
	 1
	 2 2
	 3 3 3
	 4 4 4 4

	*/
}

func loop4() {
	for i := 1; i <= 4; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(i+j, "")
		}
		fmt.Println()
	}

	/*
	   1
	   2 3
	   3 4 5
	   4 5 6 7
	*/
}

func loop5() {
	for i := 0; i <= 4; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("%c", 64+i)
		}
		fmt.Println()
	}
	/*
	   A
	   B B
	   C C C
	   D D D D
	*/
}

func loop6() {
	for row := 0; row < 6; row++ {
		for col := 0; col < 6-row; col++ {
			fmt.Print(" *")
		}
		fmt.Println()
	}
	for row := 0; row < 6; row++ {
		for col := 0; col < row+1; col++ {
			fmt.Print(" *")
		}
		fmt.Println()
	}
	/*
							 * * * * * *
		 					 * * * * *
							 * * * *
							 * * *
							 * *
							 *
		 					 *
							 * *
		                     * * *
		                     * * * *
		                     * * * * *
		                     * * * * * *


	*/
}

func loop7() {
	for i := 0; i < 3; i++ {
		for sp := 0; sp < i; sp++ {
			fmt.Print(" ")
		}
		for col := 0; col < 5-2*i; col++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 0; i < 2; i++ {
		for sp := 0; sp < 1-i; sp++ {
			fmt.Print(" ")
		}
		for col := 0; col < 2+2*i+1; col++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	/*
					*****
					 ***
		 			  *
					 ***
					*****
	*/
}
func loop8() {
	var row int
	fmt.Println("enter the row")
	fmt.Scanln(&row)
	for i := 1; i <= row; i++ {
		if i == i+1 {
			fmt.Println("*")
		} else {
			fmt.Println(i)
		}
	}
	/*
		5432*
		543*1
		43*21
		5*321
		*4321
	*/
}
func loop9() {
	var row1 int
	fmt.Println("enter the row")
	fmt.Scanln(&row1)
	for row := 0; row < row1; row++ {
		for col := 0; col < row1; col++ {
			if col == row {
				fmt.Print("*")
			} else {
				fmt.Print("0")
			}

		}
		fmt.Print("*")
		for col := row1 - 1; col >= 0; col-- {
			if col == row {
				fmt.Print("*")
			} else {
				fmt.Print("0")
			}
		}
		fmt.Println()
	}

	/*
			    *000*000*
				0*00*00*0
		    	00*0*0*00
			    000***000
	*/
}
func loop10() {
	var num int
	fmt.Println("Enter any number")
	fmt.Scanln(&num)
	for i := 1; i <= num; i++ {
		fmt.Printf("square of %d  is %d\n ", i, i*i)
	}
	fmt.Println()
}
func loop11() {
	for row := 0; row < 7; row++ { //this loop is row
		if row == 0 || row == 6 {
			for col := 0; col < 7; col++ { //this loop for colon
				fmt.Print("4 ")
			}
		} else {
			fmt.Print("4 ")
			if row == 1 || row == 5 {
				for sp := 0; sp < 5; sp++ { //this loop is space
					fmt.Print("3 ")
				}
				fmt.Print("4 ")
			} else {
				fmt.Print("3 ")
				if row == 2 || row == 4 {
					for sp := 0; sp < 4; sp++ { //this loop is space
						fmt.Print("2 ")
					}
					fmt.Print("4 ")
				} else {
					fmt.Print("2 ")
					if row == 2 || row == 4 {
						for sp := 0; sp < 4; sp++ { //this loop is space
							fmt.Print("1 ")
						}
						fmt.Print("3 ")
					} else {
						fmt.Print("1 ")
					}
				}
			}

		}
		fmt.Println()
	}

	/*
	   4 4 4 4 4 4 4
	   4 3 3 3 3 3 4
	   4 3 2 2 2 3 4
	   4 3 2 1 2 3 4
	   4 3 2 2 2 3 4
	   4 3 3 3 3 3 4
	   4 4 4 4 4 4 4
	*/

}
func loop12() {
	var num int
	fmt.Println("enter any number")
	fmt.Scanln(&num)
	for i := 1; i <= 10; i++ {
		fmt.Println(num, "*", i, "=", num*i)
	}
}
func loop13() {
	for i := 0; i < 5; i++ {
		for s := 0; s < 5-i; s++ {
			fmt.Print("")
		}
		for j := 1; j <= i; j++ {
			fmt.Print(" ", j)
		}
		fmt.Println()
	}
	/*
		    1
		   1 1
		  1 2 1
		 1 3 3 1
		1 4 6 4 1
	*/
}
func EvenAndOdd() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("even number", i)
		} else {
			fmt.Println("odd number", i)
		}

	}
}
func EvenAndOdd1() {
	i := 1
	even := 1
	for even <= 10 {
		if i%2 == 0 && even <= 10 {
			fmt.Println(i, "is even number")
			even++
		} else {
			fmt.Println(i, "is odd number")
		}
		i++
	}
	//fmt.Println("i =", i, "even number =", even)
}
func PrintUserInput() {
	var num int
	fmt.Println("enter any number")
	fmt.Scanln(&num)
	fmt.Println("table of ", num)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d*%d =%d ", num, i, num*i)
		fmt.Println()
	}

}
func pattern() {
	var num int
	fmt.Println("enter your any number")
	fmt.Scanln(&num)
	for i := 1; i <= num; i++ {
		fmt.Print("* ")
	}
}
func loop14() {
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Print((i-1)*10+j, " ")
		}
		fmt.Println()
	}
	/*
		output
		1 2 3 4 5
		11 12 13 14 15
		21 22 23 24 25
		31 32 33 34 35
		41 42 43 44 45
	*/
}
func loop15() {
	for row := 1; row <= 5; row++ {
		for col := 1; col <= 5; col++ {
			fmt.Print((row-1)*5+col, " ")
		}
		fmt.Println()
	}
	/*
		output
		1 2 3 4 5
		6 7 8 9 10
		11 12 13 14 15
		16 17 18 19 20
		21 22 23 24 25
	*/
}
func loop16() {
	num := 1
	for row := 0; row < 5; row++ {
		for sp := 1; sp <= 5-row; sp++ {
			fmt.Print(" ")
		}
		for col := 1; col <= row; col++ {
			fmt.Print(num)
			num++
		}
		fmt.Println()
	}
}
func loop17() {
	for row := 1; row <= 5; row++ {
		for sp := 1; sp <= 5-row; sp++ {
			fmt.Print(" ")
		}
		for col := 1; col <= row; col++ {
			if row%2 == 0 {
				fmt.Print("-")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
	for row := 1; row <= 5-1; row++ {
		for sp := 1; sp <= row; sp++ {
			fmt.Print(" ")
		}
		for col := 1; col <= 5-row; col++ {
			if row%2 == 0 {
				fmt.Print("*")
			} else {
				fmt.Print("-")
			}
		}
		fmt.Println()
	}
	/*
	       *
	     - -
	   * * *
	 - - - -
	   * * *
	     - -
	       *
	*/
}
func loop18() {
	for row := 0; row < 5; row++ {
		for sp := 0; sp < 2*5-row-1; sp++ {
			fmt.Print(" ")
		}
		for col := 0; col < 2*row+1; col++ {
			if col == 0 || col == 2*row || row == 5-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	/*
		     *
		   *   *
		  *     *
		 *       *
		* * * * * *
	*/
}
