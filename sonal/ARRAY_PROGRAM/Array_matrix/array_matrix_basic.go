package main

import "fmt"

/*
Add Two Matrices
Arithmetic Operations on Matrix
Determinant of a Matrix
Identity Matrix
Interchange Matrix Diagonals
Matrix Lower Triangle
Matrix Multiplication
Matrix Upper Triangle
Print Matrix Items
Scalar Matrix Multiplication
Sparse Matrix
Sum of Each Matrix Column
Sum of Each Matrix Row
Sum of Each Row and Column of a Matrix
Sum of Matrix Diagonal
Sum of Matrix Lower Triangle
Sum of Matrix Opposite Diagonal
Sum of Matrix Upper Triangle
Symmetric Matrix
Transpose a Matrix
Check Two Matrixes are Equal
*/
func main() {
	//AddTwoMatrics()
	ArithmeticOperationsonMatrix()
}
func AddTwoMatrics() {
	var addmatrics [10][10]int
	matrics1 := [3][4]int{
		{2, 3, 4, 5},
		{7, 6, 5, 4},
		{9, 8, 7, 6},
	}
	fmt.Println("this is first matrix")
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("%d\t", matrics1[i][j])
		}
	}
	matrics2 := [3][4]int{
		{9, 8, 7, 6},
		{6, 5, 4, 3},
		{8, 7, 6, 5},
	}
	fmt.Println()
	fmt.Println("this is second matrix")
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("%d\t", matrics2[i][j])
		}
	}
	fmt.Println()
	fmt.Println("addition of matrix")
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			addmatrics[i][j] = matrics1[i][j] + matrics2[i][j]
		}
	}
	fmt.Println("printing addition matrix")
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("%d\t", addmatrics[i][j])
		}
	}
	fmt.Println()
}
func ArithmeticOperationsonMatrix() {
	matrix1 := [4][4]int{
		{9, 8, 7, 6},
		{5, 4, 3, 2},
		{1, 2, 3, 4},
		{7, 9, 2, 4},
	}
	fmt.Println("this is first matrix")
	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			fmt.Printf("%d\t", matrix1[row][col])
		}
		fmt.Println()
	}
	matrix2 := [4][4]int{
		{1, 2, 3, 4},
		{2, 3, 4, 5},
		{4, 5, 6, 7},
		{4, 5, 6, 7},
	}
	fmt.Println("this is second matrix")
	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			fmt.Printf("%d\t", matrix2[row][col])
		}
		fmt.Println()
	}
	fmt.Println("\naddition", "subtraction", "multiplication", "division", "module\t")
	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			fmt.Print("\n", matrix1[row][col]+matrix2[row][col], "\t")
			fmt.Print(" ", matrix1[row][col]-matrix2[row][col], "\t")
			fmt.Print(" ", matrix1[row][col]*matrix2[row][col], "\t")
			fmt.Print(" ", matrix1[row][col]/matrix2[row][col], "\t")
			fmt.Print(" ", matrix1[row][col]%matrix2[row][col], "\t")
		}
	}
	fmt.Println()
}
