package main

import "fmt"

type student struct {
	name  string
	class string
	roll  int
}

func (so student) Student() {
	fmt.Println(so.name)
	fmt.Println(so.class)
	fmt.Println(so.roll)
}

func main() {
	s := student{}
	s.name = "sonal kumar singh"
	s.class = "BCA 3year"
	s.roll = 12
	s.Student()
	fmt.Println()
	s1 := student{}
	s1.name = "paras kumar singh"
	s1.class = "BCA final year"
	s1.roll = 13
	s1.Student()
	fmt.Println()
	s2 := student{}
	s2.name = "harshit kumar singh"
	s2.class = "mere hi sath hai"
	s2.roll = 14
	s2.Student()
	fmt.Println()
	s3 := student{}
	s3.name = "saurabh kumar singh"
	s3.class = "final year"
	s3.roll = 15
	s3.Student()
	fmt.Println()
	s4 := student{}
	s4.name = "asmirity kumari pathak"
	s4.class = "3rd year"
	s4.roll = 16
	s4.Student()
	fmt.Println()
	s5 := student{}
	s5.name = "kajal kumari"
	s5.class = "3rd year"
	s5.roll = 17
	s5.Student()
	fmt.Println()
	s6 := student{}
	s6.name = "dev kumar"
	s6.class = "3rd year"
	s6.roll = 18
	s6.Student()
	fmt.Println()
	s7 := student{}
	s7.name = "nisant oajha"
	s7.class = "3rd year"
	s7.roll = 19
	s7.Student()
	fmt.Println()
	s8 := student{}
	s8.name = "naman kumar"
	s8.class = "final year"
	s8.roll = 20
	s8.Student()
	fmt.Println()
	s9 := student{}
	s9.name = "kundan kumar singh"
	s9.class = "final year"
	s9.roll = 21
	s9.Student()

}
