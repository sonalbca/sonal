package main

import "fmt"

func main() {
	//Pattern1()
	//Pattern2()
	//Pattern3()
	//Pattern4()
	//Pattern5()
	//Pattern6()
	//Pattern7()
	//Pattern8()
	//Pattern9()
	//Pattern10()
	//Pattern11()
	pattern12() //ye bhi glat hai
	//pattern13() //ans not good
	//pattern14()//ye sahi bana hai
	//pattern15()
	//pattern16()
	//pattern17()
	//pattern18()
	//pattern19()
	//pattern20()
	//pattern21()
	//pattern22()
	//pattern23()
	//pattern24()
	//pattern25()
	//pattern26()
	//pattern27()
	//pattern28()
	//pattern29()
	//pattern30()
	//pattern31()
	//pattern32()
	//pattern33()
}

/*----------------------- 1 ------------------*/

func Pattern1() {

	for row := 0; row < 6; row++ { //this loop is row
		if row == 0 || row == 5 {
			for col := 0; col < 5; col++ { //this loop for colon
				fmt.Print("* ")
			}
		} else {
			fmt.Print("* ")
			for sp := 0; sp < 3; sp++ { //this loop is space
				fmt.Print("  ")
			}
			fmt.Print("*")

		}
		fmt.Println()
	}

	/*       * * * * *
	         *       *
		     *       *
		     *       *
		     *       *
		     * * * * *
	*/
}

/*----------------------- 2 ------------------*/

func Pattern2() {

	for row := 1; row < 5; row++ {
		for sp := 1; sp <= 5-row; sp++ {
			fmt.Print(" ")
		}
		for col := 1; col <= row; col++ {
			fmt.Printf("%d ", row)
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

/*----------------------- 3 ------------------*/

func Pattern3() {
	for row := 1; row <= 4; row++ {
		for col := 1; col <= row; col++ {
			fmt.Print(col)
		}
		fmt.Println()
	}
	/*
		1
		1 2
		1 2 3
		1 2 3 4
	*/
}

/*----------------------- 4 --------------------*/

func Pattern4() {
	for row := 1; row <= 4; row++ {
		fmt.Print(" ", row)
		for col := 1; col <= 4-row; col++ {
			fmt.Printf(" %d", row)
		}
		fmt.Println()
	}

	/*
	   1 2 3 4
	   1 2 3
	   1 2
	   1
	*/
}

/*----------------------- 5 --------------------*/

func Pattern5() {
	for row := 1; row <= 4; row++ {
		for col := 0; col < row; col++ {
			fmt.Print(row+col, " ")
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

func Pattern6() {
	for row := 1; row <= 4; row++ {
		for col := 1; col <= row; col++ {
			if (row+col)%2 == 0 {
				fmt.Print(" 1")
			} else {
				fmt.Print(" 0")
			}
		}
		fmt.Println()
	}
	/*
		     1
			 0 1
			 1 0 1
			 0 1 0 1
	*/
}

func Pattern7() {
	for row := 1; row < 4; row++ {
		for sp := 0; sp < 4-row; sp++ {
			fmt.Print(" ")
		}
		for col := 1; col < 4; col++ {
			fmt.Print(col)

		}
		fmt.Println()
	}

	/*
		      1
		    2 1 2
		  3 2 1 2 3
		4 3 2 1 2 3 4

	*/
}
func Pattern8() {
	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(" *")
		}
		fmt.Println()
	}
	/*
		    	 *
			 	 * *
			 	 * * *
			 	 * * * *
			 	 * * * * *
	*/
}
func Pattern9() {
	for row := 0; row < 7; row++ {
		for col := 0; col < 7; col++ {
			if (row+col)%2 == 0 {
				fmt.Print("1")
			} else {
				fmt.Print("0")
			}
		}
		fmt.Println()
	}
	/*
		1010101
		0101010
		1010101
		0101010
		1010101
		0101010
		1010101
	*/
}
func Pattern10() {
	for row := 0; row <= 4; row++ {
		for col := 1; col < 1+row; col++ {
			fmt.Printf("%d ", row+2)
		}
		fmt.Println()
	}
	for row := 0; row < 3; row++ {
		for col := 1; col <= 3-row; col++ {
			fmt.Printf("%d ", 5-row)
		}
		fmt.Println()
	}

	/*
		    3
			44
			555
			6666
			555
			44
			3
	*/
}
func Pattern11() {
	for i := 1; i <= 4; i++ {
		for col := 0; col < i-1; col++ {
			if i != 1 {
				fmt.Print(i)
				fmt.Print("*")
			}

		}
		fmt.Print(i)
		fmt.Println()
	}
	for i := 0; i < 4; i++ {
		for col := 1; col <= 4-i; col++ {
			fmt.Print(4 - i)
			if col != 4-i {
				fmt.Print("*")
			}
		}
		fmt.Println()

	}

	/*
			1
			2*2
			3*3*3
			4*4*4*4
		    4*4*4*4
			3*3*3
			2*2
			1
	*/
}
func pattern12() {
	for i := 1; i <= 4; i++ {
		for j := 1; j < i; j++ {
			fmt.Print(j + i)
			if j != i-1 {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
	/*
		1
		2*3
		4*5*6
		7*8*9*10
		7*8*9*10
		4*5*6
		2*3
		1
	*/
}
func pattern13() {
	for row1 := 1; row1 <= 5; row1++ {
		for col := 1; col <= row1; col++ {
			if col == 1 {
				fmt.Print("*")
			} else {
				fmt.Print(col - 1)
			}
		}
		fmt.Println()
	}
	/*
		    *
			* 1 *
			* 1 2 1 *
			* 1 2 3 2 1 *
			* 1 2 3 4 3 2 1 *
			* 1 2 3 2 1 *
			* 1 2 1 *
			* 1 *
			*
	*/
}
func pattern14() {
	for i := 0; i < 5; i++ {
		for sp := 0; sp < 5+i; sp++ {
			fmt.Print(" ")
		}
		for j := 0; j < 5; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	/*
			* * * *
			 * * * *
			  * * * *
			   * * * *
		    	* * * *
	*/
}
func pattern15() {
	/*
		   *
		  * *
		 * * *
		* * * *
		 * * *
		  * *
		   *
	*/
}
func pattern16() {
	/*
		*             *
		  *        *
		*   *   *     *
		  *   *   *
		*   *    *    *
		  *        *
		*             *

	*/
}
func pattern17() {
	/*
	* * * * * *
	* * * * * *
	* * * * * *
	* * * * * *
	* * * * * *
	* * * * * *
	 */
}
func pattern18() {
	/*
	*
	* *
	* * *
	* * * *
	* * * * *
	 */
}
func pattern19() {
	/*
	* * * * *
	* * * *
	* * *
	* *
	*
	 */
}
func pattern20() {
	/*
		        *
		      * *
		    * * *
		  * * * *
		* * * * *
	*/
}
func pattern21() {
	/*
		* * * * *
		  * * * *
		    * * *
		      * *
		        *
	*/
}
func pattern22() {
	/*
		    *
		   * *
		  * * *
		 * * * *
		* * * * *
	*/
}
func pattern23() {
	/*
		1 2 3 4
		 2 3 4
		  3 4
		   4
	*/
}
func pattern24() {
	var row int
	fmt.Println("Enter the row")
	fmt.Scanln(&row)
	for i := 0; i < row; i++ {
		for j := 1; j < row+1; j++ {
			fmt.Println(65)
		}
	}
	/*
		A
		AB
		ABC
		ABCD
	*/

}
func pattern25() {
	var num int
	fmt.Println("enter any number")
	fmt.Scanln(&num)
	for row := 1; row <= num; row++ {
		for col := 1; col <= num; col++ {
			fmt.Print(row)
		}
		fmt.Println()
	}
	/*
		output
		1 1 1 1 1
		2 2 2 2 2
		3 3 3 3 3
		4 4 4 4 4
		5 5 5 5 5
	*/
}
func pattern26() {
	var num int
	fmt.Println("enter any number")
	fmt.Scanln(&num)
	for row := 1; row <= num; row++ {
		for col := 1; col <= num; col++ {
			fmt.Print(col)
		}
		fmt.Println()
	}
	/*
		output
		1 2 3 4 5
		1 2 3 4 5
		1 2 3 4 5
		1 2 3 4 5
		1 2 3 4 5
	*/
}
func pattern27() {
	var num int
	fmt.Println("enter any number")
	fmt.Scanln(&num)
	for row := 1; row <= num; row++ {
		for col := 1; col <= num; col++ {
			fmt.Print(6 - col)
		}
		fmt.Println()
	}
	/*
		output
		5 4 3 2 1
		5 4 3 2 1
		5 4 3 2 1
		5 4 3 2 1
		5 4 3 2 1
	*/
}
func pattern28() {
	var num int
	fmt.Println("enter any number")
	fmt.Scanln(&num)
	for row := 1; row <= num; row++ {
		for col := 1; col <= num; col++ {
			fmt.Print((row-1)*1 + col)
		}
		fmt.Println()
	}
	/*
		output
		1 2 3 4 5
		2 3 4 5 6
		3 4 5 6 7
		4 5 6 7 8
		5 6 7 8 9
	*/
}
func pattern29() {
	var num int
	fmt.Println("enter any number")
	fmt.Scanln(&num)
	for row := 0; row < num; row++ {
		for col := 0; col <= row; col++ {
			fmt.Print(row + 1)
		}
		fmt.Println()
	}
	/*
		output
		1
		2 2
		3 3 3
		4 4 4 4
		5 5 5 5 5
	*/
}
func pattern30() {
	var num int
	fmt.Println("enter any number")
	fmt.Scanln(&num)
	for row := 0; row <= num; row++ {
		for col := 1; col <= row; col++ {
			fmt.Print(col)
		}
		fmt.Println()
	}
	/*
		output
		1
		1 2
		1 2 3
		1 2 3 4
		1 2 3 4 5
	*/
}
func pattern31() { //ye galat bana hai
	var num int
	fmt.Println("enter any number")
	fmt.Scanln(&num)
	n := 1
	for row := 1; row <= num; row++ {
		for col := 1; col <= row; col++ {
			fmt.Print(n, " ")
			n++
		}
		fmt.Println()
	}
	/*
		1
		2 3
		4 5 6
		7 8 9 10
		11 12 13 14
	*/
}
func pattern32() {
	var num int
	fmt.Println("enter any number")
	fmt.Scanln(&num)
	for row := num; row >= 0; row-- {
		for col := 0; col < row; col++ {
			fmt.Print(col + 1)
		}
		fmt.Println()
	}
	/*
		output
		1 2 3 4 5
		1 2 3 4
		1 2 3
		1 2
		1
	*/
}
func pattern33() {
	var n int
	n = 1
	for row := 1; row <= 5; row++ {
		for col := 1; col <= 5-row+1; col++ {
			fmt.Print(n, " ")
			n++
		}
		fmt.Println()
	}
	/*
			output
		1 2 3 4 5
		6 7 8 9
		10 11 12
		13 14
		15
	*/
}
