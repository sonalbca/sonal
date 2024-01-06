package main

import "fmt"

var a = 5
var b = 6
var c = 5
var d = 8
var e = 9

func Firsto() {
	first := (a > c && b == d || d != c)
	fmt.Println(first) //ans:=true
}
func Seco() {
	sec := (c != d || b > a && c == a)
	fmt.Println(sec) //ans:=true
}
func Thirdo() {
	third := (a != b || b != a && a != c || c != d)
	fmt.Println(third) //ans:=true
}
func Fortho() {
	forth := (b == d || d == c && c > b && c != a)
	fmt.Println(forth) //ans:=false
}
func Fiftho() {
	fift := (b == 5 && b < 4 && c > a || d != b)
	fmt.Println(fift) //ans:=true
}
func Sixtho() {
	six := (a != d && c == b || d == 9 && b < 9)
	fmt.Println(six) //ans:=false
}
func Seventho() {
	sev := !(a == b && b != c || c > a && b == c)
	fmt.Println(sev) //ans:=true
}
func Eightho() {
	eig := !(b > a && c == a) && (a != b || b == c)
	fmt.Println(eig) //ans:=false
}
func Ninetho() {
	nine := (9 == a || a == c) && (b >= 9 || d >= c)
	fmt.Println(nine) //ans:=true
}
func Tentho() {
	ten := !(a == 5) || (b >= 8)
	fmt.Println(ten) //ans:=false
}
func Eleventho() {
	elev := !(c == 9) && (b == 3) || (c == 8)
	fmt.Println(elev) //ans:=false
}
func Twelvetho() {
	twel := !(a == c) && (b > d) && (b <= e)
	fmt.Println(twel) //ans:=false
}
func Thirteentho() {
	thirt := !(a == 10) || (b >= e) && (a == 5)
	fmt.Println(thirt) //ans:=true
}
func Forteeno() {
	fort := !(a == c) && (d >= c)
	fmt.Println(fort) //ans:=false
}
func Fifteeno() {
	fift := !(e > a) || (e == 9) && (e >= 9)
	fmt.Println(fift) //ans:=true
}
func main() {
	Firsto()
	//Seco()
	//Thirdo()
	//Fortho()
	//Fiftho()
	//Sixtho()
	//Seventho()
	//Eightho()
	//Ninetho()
	//Tentho()
	//Eleventho()
	//Twelvetho()
	//Thirteentho()
	//Forteeno()
	//Fifteeno()
}
