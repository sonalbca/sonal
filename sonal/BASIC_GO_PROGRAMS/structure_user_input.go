package main

import "fmt"

func main() {
	//StrucT1()
	//StrucT2()

	//StrucT3()
	//StrucT4()
	StrucT5()
}

// ------------------------ 1 --------------------------//
type Rectangle struct {
	length float32
	breath float32
	colour string
}

func StrucT1() {
	a := Rectangle{}
	fmt.Println("enter your length")
	fmt.Scanln(&a.length)
	fmt.Println("enter your breath")
	fmt.Scanln(&a.breath)
	fmt.Println("enter your colour")
	fmt.Scanln(&a.colour)
	fmt.Println(a)
}

// --------------------------- 2 -------------------------//
type rectangle struct {
	length int
	breath int
	colour string
}

func StrucT2() {
	a := rectangle{10, 20, "red"}
	fmt.Println(a)
	fmt.Println(rectangle{23, 131, "blue"})
}

// ------------------------- 3 -----------------------------//
type REctangle struct {
	length   int
	breath   int
	color    string
	geometry struct {
		area      int
		perimeter int
	}
}

func StrucT3() {
	S := REctangle{}
	fmt.Println("enter your lenth")
	fmt.Scanln(&S.length)
	fmt.Println("enter your breath")
	fmt.Scanln(&S.breath)
	fmt.Println("enter your color")
	fmt.Scanln(&S.color)
	S.geometry.area = S.length * S.breath
	S.geometry.perimeter = 2 * (S.length * S.breath)

	fmt.Println(S)
	fmt.Println("Area", S.geometry.area)
	fmt.Println("Perimeter", S.geometry.perimeter)
}

//------------------------------ 4 --------------------------//

type REctangLe struct {
	length   int
	breath   int
	color    string
	geometry struct {
		area      int
		perimeter int
	}
}

func StrucT4() {
	S := REctangLe{
		length: 12, breath: 6, color: "sky",
	}
	S.geometry.area = S.length * S.breath
	fmt.Println("Area", S.geometry.area)

	S.geometry.perimeter = 2 * (S.length * S.breath)
	fmt.Println("Perimeter", S.geometry.perimeter)
}

//---------------------------- 5 --------------------------//

type Salary struct {
	Basic, Hra, Ta float32
}
type Employee struct {
	FirstName, LastName, Email string
	age                        int
	MonthlySalary              Salary
}

func StrucT5() {
	Emp := Employee{}
	fmt.Println("enter your first name")
	fmt.Scanln(&Emp.FirstName)
	fmt.Println("Enter your second name")
	fmt.Scanln(&Emp.LastName)
	fmt.Println("Enter your email")
	fmt.Scanln(&Emp.Email)
	fmt.Println("Enter your age")
	fmt.Scanln(&Emp.age)

	ToSal := Salary{}
	fmt.Println("Enter your Basic salary")
	fmt.Scanln(&ToSal.Basic)
	fmt.Println("Enter your Hra amount")
	fmt.Scanln(&ToSal.Hra)
	fmt.Println("Enter your Ta amount")
	fmt.Scanln(&ToSal.Ta)

	fmt.Println("First name :", Emp.FirstName)
	fmt.Println("First name :", Emp.LastName)
	fmt.Println("First name :", Emp.Email)
	fmt.Println("First name :", Emp.age)

	fmt.Println("Your basic salary is ", ToSal.Basic)
	fmt.Println("Your Hra mount is ", ToSal.Hra)
	fmt.Println("Your Ta amount is ", ToSal.Ta)
}
